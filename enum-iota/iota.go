package main

import (
	"fmt"
)

const (
	Zero = iota
	One
	Two
	Three
)

const (
	A = iota
	B
	C
)

func main() {
	fmt.Println("Value Zero:", Zero)
	fmt.Println("Value One:", One)
	fmt.Println("Value Two:", Two)
	fmt.Println("Value Three:", Three)
	fmt.Println()
	fmt.Println("Value A:", A)
	fmt.Println("Value B:", B)
	fmt.Println("Value C:", C)
}
