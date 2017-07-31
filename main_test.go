package main

import (
	"testing"
	"os"
	"github.com/beplus/makeicon/app/helper/image-helper"
	"image"
)

func TestSave(t *testing.T) {
	os.RemoveAll("AppIcon")

	save("test/icon.png")

	// check folders after test
	checkFolderExist(t, "AppIcon")
	for _, item := range image_helper.IconResized {
		checkFolderExist(t, "AppIcon"+item.Path+item.Name+".png")

		// check image size
		checkImageSize(t, "AppIcon"+item.Path+item.Name+".png", item.Width, item.Height)
	}

	os.RemoveAll("AppIcon")
}

func checkImageSize(t *testing.T, path string, width uint, height uint) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		t.Fatal(err)
	}

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		t.Fatalf("%s: %v !\n", path, err)
	}

	if (img.Width != int(width) || img.Height != int(height) || img.Width != img.Height) {
		t.Fatalf("Image: %v\n has invalid size! \n", path)
	}
}

func checkFolderExist(t *testing.T, folder string) {
	stat, err := os.Stat(folder)
	if err != nil {
		t.Fatal(err)
	}

	if stat == nil {
		t.Fatalf("%v folder was not created.", folder)
	}
}
