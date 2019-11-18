package main

import (
	"fmt"
)

type custom struct {
	str  string
	numb int
}

func main() {
	var x interface{} = 42 // x has dynamic type 'int' and value "42"
	fmt.Println("x:", x)

	x = "string" // now x has dynamic type 'string' and value "string"
	fmt.Println("x:", x)

	x = custom{"string", 42} // now x has dynamic type 'main.custom'
	fmt.Println("x:", x)
}
