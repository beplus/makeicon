// main.go
package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	"log"
	"strings"
	"bufio"
	"encoding/base64"
	"github.com/beplus/makeicon/app/helper/image-helper"
	"fmt"
)

var args struct {
	Filename string `short:"f" long:"file" description:"filename to make assets"`
}

func main() {
	_, err := flags.ParseArgs(&args, os.Args)
	if err == nil {

		if args.Filename == "" {
			log.Fatal("No file set use -n or --file to set it")
		}

		// todo windows \
		s := strings.Split(args.Filename, "/")
		filename := s[len(s)-1]
		// prepare json body

		filenameArray := strings.Split(filename, ".")
		name, extension := filenameArray[0], filenameArray[1]

		myImage, err := image_helper.NewMyImageFromBase64(getImageBase64(args.Filename), name, extension)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Processing icons... It could few seconds...")

		_, err = myImage.Upload()
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Icons saved to folder 'AppIcon'.")
		}
	}
}

func getImageBase64(filename string) string {
	imgFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("Open file fail")
	}
	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	return base64.StdEncoding.EncodeToString(buf)
}
