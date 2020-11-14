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

	"github.com/djboboch/image-convert/pkg/settings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/nickalie/go-webpbin"
)

type Flags struct {
	directoryName string
	filename      string
}

func main() {

	// Flag parsing
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

	//Creates an instance of settings
	s := settings.GetSettings()

	path, err := os.Getwd()
	if err != nil {
		log.Fatal()
	}

	s.SetCallPath(path)

	if flags.directoryName != "" {
		files, err := ioutil.ReadDir(flags.directoryName)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			if strings.Contains(f.Name(), ".jpg") {

				ConvertToWebp(filepath.Join(s.GetCallPath(), flags.directoryName, f.Name()))

			}
		}
	}

	if flags.filename != "" {
		ConvertToWebp(flags.filename)
	}

}

// ConvertToWebp converts the passed in filename .jpeg or .png to .webp file
func ConvertToWebp(path string) {
	fmt.Println(path)

	openedImage := OpenImage(path)

	defer openedImage.Close()

	decodedImage := DecodeImage(openedImage)

	filename := strings.Split(openedImage.Name(), ".")[0]

	EncodeWebp(decodedImage, filename)

}

// OpenImage converts the passed in image path to golang File type
func OpenImage(path string) *os.File {

	image, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return image

}

// DecodeImage decodes jpeg or png to Golang Image Interface
func DecodeImage(f *os.File) image.Image {

	decodedImage, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return decodedImage

}

func EncodeWebp(imageData image.Image, filename string) {

	// Function takes in a image.Image and name of file
	// Creates a new file with .webp extension to save to
	newImage, err := os.Create(filename + ".webp")
	if err != nil {
		log.Fatal(err)
	}
	// Encodes the image.Image interface into the newly created .webp file
	if err := webpbin.Encode(newImage, imageData); err != nil {
		newImage.Close()
		log.Fatal(err)
	}
	// closes the files open stream
	if err := newImage.Close(); err != nil {
		log.Fatal(err)
	}
}
