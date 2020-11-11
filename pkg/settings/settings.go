package settings

type Settings interface {
	SetCallPath()
	GetCallPath() string
}

type settings struct {
	callpath string
}

var instance *settings

func GetSettings() *settings {
	if instance == nil {
		instance = new(settings)
	}

	return instance
}

func (s *settings) SetCallPath(path string) {
	s.callpath = path
}

func (s *settings) GetCallPath() string {

	return ""
}
