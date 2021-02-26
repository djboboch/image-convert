package folder

import "github.com/djboboch/image-convert/models/file"

// Folder represents the stored data of the folder
type Folder struct {
	Name    string
	Files   []file.File
	Folders []*Folder
}

func New(name string) *Folder {
	return &Folder{name, []file.File{}, []*Folder{}}
}
