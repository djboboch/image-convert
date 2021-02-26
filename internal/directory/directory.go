package directory

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	fileModel "github.com/djboboch/image-convert/models/file"
	"github.com/djboboch/image-convert/models/folder"
)

// Read directory and return slice of Fileinfo
func read(directory string) []os.FileInfo {

	files, err := ioutil.ReadDir(directory)
	
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func CreateTree(path string) *folder.Folder {

	rootFolder := folder.New(path)

	for _, file := range read(path) {

		fileExtenstion := filepath.Ext(file.Name())

		if file.IsDir() {
			// Call CreateFolderTree for the inside folder
			rootFolder.Folders = append(rootFolder.Folders, CreateTree(filepath.Join(path, file.Name())))
		} else if fileExtenstion == ".jpg" || fileExtenstion == ".png" {
			rootFolder.Files = append(rootFolder.Files, fileModel.New(filepath.Join(path, file.Name())))
		}
	}

	return rootFolder
}