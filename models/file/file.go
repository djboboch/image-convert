package file

// File represents the stored data of the file
type File struct {
	Name string
}

func New(name string) File {
	return File{name}
}