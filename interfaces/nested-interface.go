package main

import (
	"fmt"
	"strings"
)

type Interface1 interface {
	Method1(int, string) (int, error)
}
type Interface2 interface {
	Method2(string)
	Interface1 //Embedding Interface1
}

type Interface3 interface {
	Method3(string) bool
	Interface2 //Embedding Interface2
}

type T struct{}

func (T) Method1(val int, str string) (int, error) {
	return 42, nil
}

func (T) Method2(str string) {
	fmt.Println(str)
}

func (T) Method3(str string) bool {
	return strings.Contains(str, "yes")
}

func main() {
	i := T{}
	i.Method1(0, "0")
	i.Method2("Hello world!")
	i.Method3("No")
}
