package main

import (
	"fmt"
)

func main() {
	b := make([]int, 5)
	fmt.Println("step-1", len(b), cap(b), b)

	b = make([]int, 5, 20)
	fmt.Println("step-2", len(b), cap(b), b)

	sentinel := 5
	for i := 0; i < 20; i++ {
		if i != 0 && i%sentinel == 0 {
			b = b[:len(b)+sentinel]
			fmt.Println("loop", i, len(b), cap(b), b)
		}
		b[i] = i
	}
	fmt.Println("step-3", len(b), cap(b), b)

	b = b[:cap(b)]
	fmt.Println("step-4", len(b), cap(b), b)
}
