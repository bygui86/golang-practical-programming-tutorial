package main

import (
	"fmt"
	"strings"
)

type Interface1 interface {
	Method1(int, string) (int, error)
}
type Interface2 interface {
	Interface1 // Embedding Interface1
	Method2(string)
}

type Interface3 interface {
	Interface2 // Embedding Interface2
	Method3(string) bool
}

type Embedded struct{
	// INFO: we can explicit specify which interfaces are implemented
	// Interface3
}

func (e *Embedded) Method1(val int, str string) (int, error) {
	return 42, nil
}

func (e *Embedded) Method2(str string) {
	fmt.Println(str)
}

func (e *Embedded) Method3(str string) bool {
	return strings.Contains(str, "yes")
}

func main() {
	var i Interface3
	i = &Embedded{}
	i.Method1(0, "0")
	i.Method2("Hello world!")
	i.Method3("No")
}
