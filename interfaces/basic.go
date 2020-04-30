package main

import (
	"fmt"
	"math"
)

// INTERFACES

type Abser interface {
	Abs() float64
}

// STRUCTS

type MyFloat float64

// MyFloat implements interface Abser directly
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// Vertex implements interface Abser through pointer (*Vertex), so we have to pay attention on how to instantiate it
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// structs
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	
	// interfaces
	var a Abser
	a = f  // a MyFloat implements Abser
	fmt.Println(a.Abs())
	
	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())
	
	a = v // v is a Vertex (not *Vertex) and does NOT implement Abser.
	fmt.Println(a.Abs())
}
