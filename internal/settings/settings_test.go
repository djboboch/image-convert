package settings

import "testing"

func TestGetSettings(t *testing.T) {
	settings := GetSettings()

	if settings == nil {
		t.Error("expected pointer to settings after callingn GetSettings(), not nil")
	}

	expectedSettings := settings

	anotherSettings := GetSettings()

	if anotherSettings != expectedSettings {
		t.Error("Expected same instance of settings in anotherSettings but got a different instance.")
	}
}

func TestSetCallPath(t *testing.T) {

	settings := GetSettings()

	settings.SetCallPath()
}
