package main

import "log"

// Interface
type Interface interface {
	Method(string) string
}

// Struct to implement the interface
type T struct{}

// Interface methods implementation on the declared Struct
func (t T) Method(str string) string {
	return str
}

func logMsg(i Interface) {
	log.Println(i.Method("Hello"))
}

func main() {
	logMsg(T{})
}
