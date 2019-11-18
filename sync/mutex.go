package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}
var accessCount int

/*
	A Mutex is a shared lock that you can use to provide exclusive access to certain parts of your code.
	In this simple example, we have a shared/global variable accessCount which is used in the incr function.

	If you run this, you will always get the same result, i.e. Final = 500
*/
func main() {

	var wg sync.WaitGroup
	loop := 500
	for i := 1; i <= loop; i++ {
		go func(c int) {
			wg.Add(1)
			defer wg.Done()
			fmt.Printf("Increment by goroutine %v\n", c)
			incr()
		}(i)
	}
}

func incr() {

	/*
		Comment (or remove) the following lines in the incr function and run the program on your local machine and run the program again
		You will notice variable results e.g. Final = 474.
	*/
	mutex.Lock()
	defer mutex.Unlock()

	accessCount = accessCount + 1
}
