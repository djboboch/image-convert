package config

type Config interface {
	SetCallPath()
	GetCallPath() string
	GetRecursiveReference() *bool
}

type config struct {
	isRecursive bool
	position    string
}

var instance *config

func GetConfig() *config {
	if instance == nil {
		instance = new(config)
	}

	return instance
}

func (c *config) SetCallPath(path string) {
	//Todo: Check if
	c.position = path
}

func (c *config) GetRecursiveReference() *bool {
	return &instance.isRecursive
}

func (s *config) GetCallPath() string {

	return s.position
}
