package main

import "log"

type Interface interface {
	Method1(int, string) (int, error)
	Method2(string)
}

type T1 struct{}

func (T1) Method1(val int, str string) (int, error) {
	return 42, nil
}

func (T1) Method2(str string) {
	log.Printf("T1 %s", str)
}

type T2 struct{}

func (T2) Method1(val int, str string) (int, error) {
	return 24, nil
}

func (T2) Method2(str string) {
	log.Printf("T2 %s", str)
}

func main() {
	var i Interface

	log.Println("Sample 1")
	i = T1{}
	i.Method1(0, "0")
	i.Method2("Hello world!")

	log.Println("Sample 2")
	i = T2{}
	i.Method1(1, "1")
	i.Method2("Hello Matt")
}
