package main

import (
	"fmt"
	"io"
)

type I interface {
	m()
}

/*
	A type assertion can be used to:
		. obtain a value of concrete type from a value of interface type
		. obtain a value of a different interface type than the initial one
		  (an interface with a different method set, practically not subset of the original one as
		  that could simply be obtained using a simple type conversion).
*/
// See https://stackoverflow.com/questions/31379404/is-this-casting-in-golang
func main() {
	var x interface{} = 7 // x has dynamic type int and value 7
	i := x.(int)          // i has type int and value 7
	fmt.Println("x:", x)
	fmt.Println("i:", i)
}

func f(y I) {
	// s := y.(string)    // illegal: string does not implement I (missing method m)
	r := y.(io.Reader) // r has type io.Reader and the dynamic type of y must implement both I and io.Reader
	fmt.Println("r:", r)
}
