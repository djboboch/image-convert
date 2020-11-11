package main

import (
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/nickalie/go-webpbin"
)

type Flags struct {
	directoryName string
	filename      string
}

type Settings struct {
}

func main() {

	flags := Flags{}
	flag.StringVar(&flags.directoryName, "dir", "", "Directory to run the conversion for")
	flag.StringVar(&flags.filename, "f", "", "File to convert")
	//recursive := flag.Bool("r", false, "Should the tool do a recursive convert.")
	flag.Parse()

	if flags.directoryName == "" && flags.filename == "" {
		fmt.Println("No -dir or -f argument provided.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if flags.directoryName != "" && flags.filename != "" {
		fmt.Println("Provide only one argument for the conversion location. Cannot use both -dir or -f.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	s := Settings{}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal()
	}
	s.callpath = path

	if flags.directoryName != "" {
		files, err := ioutil.ReadDir(flags.directoryName)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			if strings.Contains(f.Name(), ".jpg") {

				EncodeToWebp(filepath.Join(s.callpath, flags.directoryName, f.Name()))

			}
		}
	}

	if flags.filename != "" {
		EncodeToWebp(flags.filename)
	}

}

// EncodeToWebp encodes the passed
func EncodeToWebp(path string) {
	fmt.Println(path)

	openedImage, err := os.Open(path)
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
