package main

import (
	"fmt"
	"time"
)

func main() {
	n := 3
	multiCh := make(chan int)

	// We supply 1 channel to the `multiplyByTwo` function
	// for both sending and receiving data
	go multiplyByTwo(multiCh)

	// We then send it data through the channel and wait for the result
	time.Sleep(3 * time.Second)
	multiCh <- n
	fmt.Println(n, "multiplied by 2:", <-multiCh)
}

func multiplyByTwo(multiCh chan int) {
	// This line is just to illustrate that there is code that is
	// executed before we have to wait on the `in` channel
	fmt.Println("Waiting for input to multiply...")

	// The goroutine does not proceed until data is received on the `in` channel
	num := <-multiCh

	// The rest is unchanged
	result := num * 2
	multiCh <- result
}
