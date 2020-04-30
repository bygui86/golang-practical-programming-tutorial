package main

import "log"

type ExampleInterface interface {
	Method(string) string
}

type ExampleStruct struct{}

func (e ExampleStruct) Method(str string) string {
	return str
}

func main() {
	exampleStruct := ExampleStruct{}
	logMsg(exampleStruct)
	
	var exampleInterface ExampleInterface
	exampleInterface = ExampleStruct{}
	logMsg(exampleInterface)
}

func logMsg(i ExampleInterface) {
	log.Println(i.Method("Hello"))
}
