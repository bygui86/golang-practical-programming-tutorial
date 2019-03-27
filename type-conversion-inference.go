package main

import "fmt"

func intToFloat() {
	var a int = 62
	var b float64 = float64(a)
	fmt.Println("a and b:", a, b)
}

func typeInference() {
	var x float32
	y := x // y is float32 type
	fmt.Println("x and y:", x, y)
}

func main() {
	intToFloat()
	typeInference()
}
