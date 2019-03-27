package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func typeValue() {
	fmt.Println("Type: %T Value: %v", ToBe, ToBe)
	fmt.Println("Type: %T Value: %v", MaxInt, MaxInt)
	fmt.Println("Type: %T Value: %v", z, z)
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func typeLiterals() {
	fmt.Println(v1, p, v2, v3)
	fmt.Printf("%T - %v \n", v1, v1)
	fmt.Printf("%T - %v \n", p, p)
}

func main() {
	typeValue()
	fmt.Println()
	typeLiterals()
}
