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

var version = "master"

var args struct {
	Filename string `short:"f" long:"file" description:"File path to make the Icon from. Have to be square, 1024px x 1024px at least."`
	Version  bool `short:"v" long:"version" description:"Show version."`
}

func main() {
	_, err := flags.ParseArgs(&args, os.Args)
	if err == nil {
		if args.Version {
			fmt.Printf(" Version: %v \n License: MIT \n Copyright (c) 2017 BePlus s.r.o. (https://be.plus/makeicon) \n", version)
		} else {

			if args.Filename == "" {
				log.Fatal("No file specified. Provide the path to the file using -f or --file.")
			}

			save(args.Filename)
		}
	}
}

func save(file string) {
	// todo windows \
	s := strings.Split(file, "/")
	filename := s[len(s)-1]

	// prepare json body
	filenameArray := strings.Split(filename, ".")
	name, extension := filenameArray[0], filenameArray[1]

	myImage, err := image_helper.NewMyImageFromBase64(getImageBase64(args.Filename), name, extension)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Processing your Icons... It may take a while.")

	_, err = myImage.Upload()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Your Icons was saved to a folder 'AppIcon'.")
	}
}

func getImageBase64(filename string) string {
	imgFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file.")
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
