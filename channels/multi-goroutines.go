package main

import (
	"fmt"
	"time"
)

func main() {
	out := make(chan int)
	in := make(chan int)

	// Create 3 `multiplyByTwo` goroutines.
	/* WARNING
	It is important to note that there is no guarantee as to which goroutine will accept which input,
	or which goroutine will return an output first. All the main function “knows”, is that it is sending
	some data into the in channel, and expects some data to be received on the out channel.
	*/
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)

	// Up till this point, none of the created goroutines actually do
	// anything, since they are all waiting for the `in` channel to
	// receive some data
	time.Sleep(3 * time.Second)
	in <- 1
	in <- 2
	in <- 3

	// Now we wait for each result to come in
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

func multiplyByTwo(in <-chan int, out chan<- int) {
	fmt.Println("Waiting for input to multiply...")
	num := <-in
	result := num * 2
	out <- result
}
