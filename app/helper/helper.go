package helper

import (
	"regexp"
	"strings"
	"time"
)

func StrToTime(datetime string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", datetime, time.Local)
}

func NowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Превратить русскоформатный номер в английский
func ClearPlate(plate string) string {
	out := strings.ToUpper(plate)
	replacer := strings.NewReplacer("А", "A", "В", "B", "Е", "E", "К", "K", "М", "M", "Н", "H", "О", "O", "Р", "P", "С", "C", "Т", "T", "У", "Y", "Х", "X")
	out = replacer.Replace(out)
	regex := regexp.MustCompile("[^A-Z0-9]")
	out = regex.ReplaceAllString(out, "")
	return out
}

func InArray(arr *[]string, str string) bool {
	if arr == nil {
		return false
	}
	for _, a := range *arr {
	   if a == str {
		  return true
	   }
	}
	return false
 }

 // AddSlashes если тупит GORM
func AddSlashes(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `'`, `\'`)
	return s
}