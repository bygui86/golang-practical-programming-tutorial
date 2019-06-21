package main

import "log"

type Interface interface {
	Method(string) string
}

type T struct{}

func (T) Method(str string) string {
	return str
}

func logMsg(i Interface) {
	log.Println(i.Method("Hello"))
}

func main() {
	logMsg(T{})
}
