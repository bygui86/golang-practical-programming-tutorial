package main

import (
	"fmt"
	"math"
)

func regular() {
	x := 5
	if x > 25 {
		fmt.Println("Smaller!")
	} else {
		fmt.Println("Bigger!")
	}
}

func shortStatement() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	regular()
	shortStatement()
}
