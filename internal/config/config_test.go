package config

import (
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	settings := GetConfig()

	if settings == nil {
		t.Error("expected pointer to settings after callingn GetSettings(), not nil")
	}

	expectedSettings := settings

	anotherSettings := GetConfig()

	if anotherSettings != expectedSettings {
		t.Error("Expected same instance of settings in anotherSettings but got a different instance.")
	}
}

func TestConfig_GetRecursiveReference(t *testing.T) {

	config := GetConfig()

	got := config.GetRecursiveReference()

	var want *bool

	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("Type of got is: %s it should be *bool", reflect.TypeOf(got))
	}

}
