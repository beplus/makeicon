package image_helper

import (
	"encoding/base64"
	"bytes"
	"github.com/nfnt/resize"
	"image/png"
	"image"
	"image/jpeg"
	"os"
	"errors"
	"github.com/mholt/archiver"
	"time"
	"io/ioutil"
	"github.com/beplus/makeicon/app/helper/s3-helper"
	"strconv"
	"github.com/disintegration/imaging"
)

type MyImage struct {
	Image image.Image
	File  File
}

type File struct {
	Name      string
	Extension string
}

type ResizeImage struct {
	Path   string
	Name   string
	Width  uint
	Height uint
}

func NewMyImageFromBase64(base64String string, name string, extension string) (*MyImage, error) {
	image, err := getImageFromBase64(base64String, extension)
	if err != nil {
		return nil, err
	}

	return &MyImage{Image: image, File:File{Name:name, Extension:extension}}, nil
}

func getImageFromBase64(base64String string, extension string) (image.Image, error) {
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(data)

	var img image.Image

	switch extension {
	case "png":
		img, err = png.Decode(r)
	//case "jpg":
	//	img, err = jpeg.Decode(r)
	default:
		return img, errors.New("unknown file extension " + extension)
	}

	return img, err
}

func (m MyImage) checkIconProperties() (error) {
	x, y := m.Image.Bounds().Dx(), m.Image.Bounds().Dy()
	if x != y {
		return errors.New("Image is not square")
	}
	if x < 1024 {
		return errors.New("Image resollution is smaler then 1024")
	}

	return nil
}

func (m MyImage) checkSplashProperties() (error) {
	x, y := m.Image.Bounds().Dx(), m.Image.Bounds().Dy()
	if x > 1024 && y < 1024 {
		return errors.New("Image resollution is smaler then 1024")
	}

	return nil
}

/**
source: splash|icons
target: S3|local
 */
func (m MyImage) Upload(source string, target string, orientation string, method string) (string, error) {
	resized := []ResizeImage{}
	path := "files"
	switch source{
	case "splash":
		err := m.checkSplashProperties()
		if err != nil {
			return "", err
		}
		resized = SplashResized
		path = "files/splash"
	case "icons":
		orientation = ""
		err := m.checkIconProperties()
		if err != nil {
			return "", err
		}
		resized = IconResized
		path = "files/icons"
	default:
		return "", errors.New("Unknown source : " + source)
	}

	switch method {
	case "folder":
		return "", m.save("AppIcon", resized, orientation)
	case "zip":
		prefix := strconv.FormatInt(time.Now().UnixNano(), 10) + "_" + m.File.Name

		// create temp dir without prefix in default directory
		tempDirPath, err := ioutil.TempDir("", "")
		if err != nil {
			return "", err
		}
		defer os.RemoveAll(tempDirPath)

		err = m.save(tempDirPath, resized, orientation)
		if err != nil {
			return "", err
		}

		zipFileName := prefix + "_" + source + ".zip"

		switch target {
		case "S3":
			zipPath := tempDirPath + "/" + zipFileName

			err = archiver.Zip.Make(zipPath, []string{tempDirPath + "/" + Ios, tempDirPath + "/" + Android})
			if err != nil {
				return "", err
			}

			return s3_helper.UploadToS3(zipPath, path + m.File.Name + "/" + zipFileName)
		case "local":
			zipPath := zipFileName

			err = archiver.Zip.Make(zipPath, []string{tempDirPath + "/" + Ios, tempDirPath + "/" + Android})
			if err != nil {
				return "", err
			}

			return "", nil
		default:
			return "", errors.New("Unknown target : " + target)
		}
	default:
		return "", errors.New("Unknown method : " + method)
	}
}

func (m MyImage) save(dir string, resized []ResizeImage, orientation string) (error) {
	for _, item := range resized {
		switch orientation {
		case "portrait":
			if item.Width > item.Height {
				continue
			}
		case "landscape":
			if item.Width < item.Height {
				continue
			}
		}

		// create directories if not exist
		if _, err := os.Stat(dir + item.Path); os.IsNotExist(err) {
			os.MkdirAll(dir + item.Path, 0777)
		}
		// save image in specific size
		err := m.SaveInSize(dir + item.Path + item.Name, item.Width, item.Height)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m MyImage) SaveInSize(path string, width uint, height uint) (error) {
	x, y := m.Image.Bounds().Dx(), m.Image.Bounds().Dy()

	inputRatio := float64(x) / float64(y)
	outputRatio := float64(width) / float64(height)

	var resizedImage image.Image
	switch {
	case inputRatio < outputRatio:
		tmpImg := resize.Resize(width, 0, m.Image, resize.Lanczos2)
		resizedImage = imaging.CropCenter(tmpImg, int(width), int(height))
	case inputRatio > outputRatio:
		tmpImg := resize.Resize(0, height, m.Image, resize.Lanczos2)
		resizedImage = imaging.CropCenter(tmpImg, int(width), int(height))
	default :
		resizedImage = resize.Resize(width, height, m.Image, resize.Lanczos2)
	}

	return save(path, resizedImage, m.File.Extension)
}

func save(path string, image image.Image, extension string) (error) {
	out, err := os.Create(path + "." + extension)
	if err != nil {
		return err
	}
	defer out.Close()

	switch extension {
	case "png":
		png.Encode(out, image)
	case "jpg":
		jpeg.Encode(out, image, nil)
	default:
		return errors.New("unknown file extension " + extension)
	}
	return nil
}
