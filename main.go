package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/djboboch/image-convert/internal/directory"
	"github.com/djboboch/image-convert/internal/settings"
	"github.com/djboboch/image-convert/models/folder"

	_ "image/jpeg"
	_ "image/png"

	"github.com/nickalie/go-webpbin"
)

const ()

// Flags is a structure to represent passed in flas from the cli
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

	callPath, err := os.Getwd()
	if err != nil {
		log.Fatal()
	}

	s.SetCallPath(callPath)

	if flags.directoryName != "" {

		folderTree := directory.CreateTree(filepath.Join(s.GetCallPath(), flags.directoryName))

		folderTree.convertImages()

	}
}

// ConvertToWebp converts the passed in filename .jpeg or .png to .webp file
// TODO: Move to package
func ConvertToWebp(path string) {
	os.Chdir(path)

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

// EncodeWebp encodes the passed in image.Image data into the passed in filename string
// TODO: Move to own package
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

// Move this to internal package
// TODO: Move to internal package
func (f *folder.Folder) convertImages() {
	for _, file := range f.Files {
		ConvertToWebp(file.Name)
	}

	for _, folder := range f.Folders {
		folder.convertImages()
	}
}
