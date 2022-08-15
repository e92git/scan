package main

import (
	"fmt"
	"regexp"
	"strings"
)

const refString = "Н832МО124"

// для тестирования способностей языка го. Для запуска:
// > go run test/test.go 
func main() {
	out := strings.ToUpper(refString)
	replacer := strings.NewReplacer("А", "A", "В", "B", "Е", "E", "К", "K", "М", "M", "Н", "H", "О", "O", "Р", "P", "С", "C", "Т", "T", "У", "Y", "Х", "X")
	out = replacer.Replace(out)
	regex := regexp.MustCompile("[^A-Z0-9]")
	out = regex.ReplaceAllString(out, "")
	fmt.Println(out)
}
