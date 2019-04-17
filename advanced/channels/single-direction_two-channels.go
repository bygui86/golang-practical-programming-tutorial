package main

import (
	"fmt"
	"time"
)

func main() {
	n := 3
	in := make(chan int)
	out := make(chan int)

	// We supply 2 channels to the `multiplyByTwo` function
	// One for sending data and one for receiving
	go multiplyByTwo(in, out)

	// We then send it data through the channel and wait for the result
	time.Sleep(3 * time.Second)
	in <- n
	fmt.Println(n, "multiplied by 2:", <-out)
}

func multiplyByTwo(in <-chan int, out chan<- int) {
	// This line is just to illustrate that there is code that is
	// executed before we have to wait on the `in` channel
	fmt.Println("Waiting for input to multiply...")

	// The goroutine does not proceed until data is received on the `in` channel
	num := <-in
	result := num * 2
	out <- result
}
