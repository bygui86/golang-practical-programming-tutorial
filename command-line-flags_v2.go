package main

import (
	"flag"
	"fmt"
)

// Sarma configuration options
var (
	word    = ""    // required
	numb    = -1    // required
	boolean = false // optional
)

func init() {
	flag.StringVar(&word, "word", "", "a string")
	flag.IntVar(&numb, "numb", -1, "a number")
	flag.BoolVar(&boolean, "boolean", false, "a boolean")
	flag.Parse()

	if len(word) == 0 {
		panic("no word defined, please set the -word flag")
	}

	if numb == -1 {
		panic("no number defined, please set the -numb flag")
	}
}

func main() {
	fmt.Println("word:", word)
	fmt.Println("numb:", numb)
	fmt.Println("boolean:", boolean)
	fmt.Println("tail:", flag.Args())
}
