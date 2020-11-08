package main

import (
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/nickalie/go-webpbin"
)

func main() {

	var directory string
	var filename string
	flag.StringVar(&directory, "dir", "", "Directory to run the conversion for")
	flag.StringVar(&filename, "f", "", "File to convert")
	//recursive := flag.Bool("r", false, "Should the tool do a recursive convert.")
	flag.Parse()

	if directory == "" && filename == "" {
		fmt.Println("No -dir or -f argument provided.")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if directory != "" && filename != "" {
		fmt.Println("Provide only one argument for the conversion location. Cannot use both -dir or -f.")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if directory != "" {
		files, err := ioutil.ReadDir(directory)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			fmt.Println(f.Name())
		}
	}

	if filename != "" {
		EncodeToWebp(filename)
	}

}

// EncodeToWebp encodes the passed
func EncodeToWebp(name string) {

	openedImage, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer openedImage.Close()

	filename := strings.Split(openedImage.Name(), ".")[0]

	// Cretes the new file for conversion into webp
	newImage, err := os.Create(filename + ".webp")
	if err != nil {
		log.Fatal(err)
	}

	decodedImage := DecodeImage(openedImage)

	// Encodes the image interface into the webp format
	if err := webpbin.Encode(newImage, decodedImage); err != nil {
		newImage.Close()
		log.Fatal(err)
	}

	if err := newImage.Close(); err != nil {
		log.Fatal(err)
	}

}

// DecodeImage decodes jpeg or png to Golang Image Interface
func DecodeImage(f *os.File) image.Image {

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return img

}
