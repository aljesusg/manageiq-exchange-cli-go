package utils

import (
	"reflect"
	"testing"
)

func TestPrintColor(t *testing.T) {

	want := "\033[0;0;31mHello World\033[0m"
	if !reflect.DeepEqual(PrintColor("Hello World", "Red"), want) {
		t.Errorf("PrintColor returned %+v, want %+v", PrintColor("Hello World", "Red"), want)
	}
}

type SampleInterface struct {
	Type          string `json:"type"`
	Enabled       bool   `json:"enabled"`
	ApplicationId string `json:"id_application"`
	Server        string `json:"server"`
	Version       string `json:"version"`
	Verify        bool   `json:"verify"`
}

func TestCreateFromMap(t *testing.T) {

	want := SampleInterface{
		Server:        "github.com",
		Version:       "v3",
		ApplicationId: "abc",
	}
	var data = map[string]interface{}{
		"server":         "github.com",
		"version":        "v3",
		"id_application": "abc",
	}
	var provider SampleInterface
	CreateFromMap(data, &provider)
	if !reflect.DeepEqual(provider, want) {
		t.Errorf("CreateFromMap returned %+v, want %+v", provider, want)
	}
}
