package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
    "github.com/thoas/go-funk"
)

const refString = "Н832МО124"

// для тестирования способностей языка го. Для запуска:
// > go run test/test.go
func main2() {
	// разбитие на bulk массивы
	t := []int{1, 2, 3, 4, 5, 6, 7, 8}
	tt := funk.Chunk(t, 3)
	tt2 := tt.([][]int)
	println("chunk size =", len(tt2), "; elem = ", tt2[2][1])

	// строковые и регулярные выражения
	out := strings.ToUpper(refString)
	replacer := strings.NewReplacer("А", "A", "В", "B", "Е", "E", "К", "K", "М", "M", "Н", "H", "О", "O", "Р", "P", "С", "C", "Т", "T", "У", "Y", "Х", "X")
	out = replacer.Replace(out)
	regex := regexp.MustCompile("[^A-Z0-9]")
	out = regex.ReplaceAllString(out, "")
	fmt.Println(out)

	// джинерики
	fmt.Println(
		New(func(a int, b int) int { return a + b }),
		Max[int64](2, 34),
	)
	m := myMap[int, string]{5: "foo"}
	println(m[5]) // foo

	// канали и горутины
	c1, c2 := make(chan string), make(chan string)
	defer func() { close(c1); close(c2) }() // не забываем прибраться

	go func(c chan<- string) { println("chanrrr"); time.Sleep(time.Second); println("chanrrr2"); c <- "foo" }(c1)
	go func(c chan<- string) { <-time.After(time.Second); c <- "bar" }(c2)

	for i := 1; ; i++ {
		select { // блокируемся, пока в один из каналов не попадёт сообщение
		case val := <-c1:
			println("channel 1", val)

		case val := <-c2:
			println("channel 2", val)
		}

		if i >= 2 { // через 2 итерации выходим (иначе будет deadlock)
			break
		}
	}
}

func New(cmp func(int, int) int) int {
	return cmp(1, 2)
	// return &Tree(E){compare: cmp}
}

// Generics
type Int interface {
	~int | ~int32 | ~int64 | ~string
}

func Max[T Int](a T, b T) T {
	if a > b {
		return a
	}

	return b
}

type myMap[K comparable, V any] map[K]V
