package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const NC string = "\033[0m" // No Color

var COLOR = map[string]string{
	"Black":        "0;30",
	"Dark Gray":    "1;30",
	"Red":          "0;31",
	"Light Red":    "1;31",
	"Green":        "0;32",
	"Light Green":  "1;32",
	"Brown/Orange": "0;33",
	"Yellow":       "1;33",
	"Blue":         "0;34",
	"Light Blue":   "1;34",
	"Purple":       "0;35",
	"Light Purple": "1;35",
	"Cyan":         "0;36",
	"Light Cyan":   "1;36",
	"Light Gray":   "0;37",
	"White":        "1;37",
}

func CreateFromMap(m map[string]interface{}, result interface{}) error {
	data, _ := json.Marshal(m)
	err := json.Unmarshal(data, &result)
	if err != nil {
		return err
	}
	return nil
}

func PrintColor(phrase string, color string) string {
	return fmt.Sprintf("\033[0;%sm%s%s", COLOR[color], phrase, NC)
}

func PrintValues(object interface{}, path string, blocked []string) string {
	var result string
	s := reflect.ValueOf(object).Elem()
	typeOfT := s.Type()
	result = ""
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if stringInSlice(typeOfT.Field(i).Name, blocked) || valueIsEmpty(f.Type().String(), f.Interface()) {
			continue
		}
		result += fmt.Sprintf("%s%s : %v\n", path,
			typeOfT.Field(i).Name, f.Interface())
	}
	return result
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func valueIsEmpty(typ string, a interface{}) bool {
	switch typ {
	case "string":
		if len(a.(string)) == 0 {
			return true
		}
	}
	return false
}
