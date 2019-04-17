package main

import (
	"fmt"
	"time"
)

func main() {

	intCh := make(chan int)
	boolCh := make(chan bool)
	stringCh := make(chan string)

	go multiplexer(intCh, boolCh, stringCh)
	go multiplexer(intCh, boolCh, stringCh)

	time.Sleep(1 * time.Second)
	intCh <- 3
	time.Sleep(1 * time.Second)
	boolCh <- true

	fmt.Println(<-stringCh)
	fmt.Println(<-stringCh)
}

func multiplexer(intCh <-chan int, boolCh <-chan bool, stringCh chan<- string) {

	select {
	case n := <-intCh:
		fmt.Println("Received a number:", n)
	case b := <-boolCh:
		fmt.Println("Received a bool:", b)
	}

	stringCh <- "done"
}
