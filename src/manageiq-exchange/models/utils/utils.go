package utils

import (
	"encoding/json"
	"fmt"
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
