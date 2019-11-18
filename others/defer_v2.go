package main

import (
	"fmt"
)

// Stacking defers
// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
func main() {
	fmt.Println("start")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("end")
}
