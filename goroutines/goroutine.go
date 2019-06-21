package main

import (
	"fmt"
	"time"
)

func sample(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(500 * time.Millisecond)
	}
}

// A goroutine is a lightweight thread of execution.
func main() {

	// Suppose we have a function call f(s). Here’s how we’d call that in the usual way,
	// running it synchronously.
	sample("direct")

	// To invoke this function in a goroutine, use go f(s). This new goroutine will execute
	// concurrently with the calling one.
	go sample("goroutine")

	// You can also start a goroutine for an anonymous function call.
	go func(from string) {
		for i := 0; i < 3; i++ {
			fmt.Println(from, ":", i)
			time.Sleep(500 * time.Millisecond)
		}
	}("goroutine-inline")

	// Our two function calls are running asynchronously in separate goroutines now, so execution
	// falls through to here. This Scanln requires we press a key before the program exits.
	time.Sleep(3 * time.Second)
	fmt.Println("Press any key when ready...")
	fmt.Scanln()
	fmt.Println("done")
}
