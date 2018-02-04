package info

import (
	"fmt"
	utils "manageiq-exchange/models/utils"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	want := Info{
		Version: "1.0",
		Providers: map[string]Provider{
			"github.com": Provider{
				ApplicationId: "abc",
				Server:        "github.com",
				Version:       "v3",
			},
		},
	}
	var data = map[string]interface{}{
		"version": "1.0",
		"providers": map[string]interface{}{
			"github.com": map[string]interface{}{
				"server":         "github.com",
				"version":        "v3",
				"id_application": "abc",
			},
		},
	}
	var info Info
	info.Init(data)
	if !reflect.DeepEqual(info, want) {
		t.Errorf("InfoInit returned %+v, want %+v", info, want)
	}
}

func TestPrint(t *testing.T) {
	want := fmt.Sprintf("%s: 1.0\n\n", utils.PrintColor("Version", "Red"))
	want += fmt.Sprintf("Providers: \n")
	want += fmt.Sprintf("    github.com: \n")
	want += fmt.Sprintf("        Server: github.com\n")
	want += fmt.Sprintf("        ApplicationId: abc\n")
	want += fmt.Sprintf("        Version: v3\n")

	info := Info{
		Version: "1.0",
		Providers: map[string]Provider{
			"github.com": Provider{
				ApplicationId: "abc",
				Server:        "github.com",
				Version:       "v3",
			},
		},
	}
	if !reflect.DeepEqual(info.Print(), want) {
		t.Errorf("Info Print returned %+v, want %+v", info.Print(), want)
	}
}
