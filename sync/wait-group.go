package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

/*
	Use a WaitGroup for co-ordination if your program needs to wait for a bunch of Goroutines to finish.
	It is similar to a CountDownLatch in Java.
*/
func main() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	filesInHomeDir, err := ioutil.ReadDir(homeDir)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(filesInHomeDir))
	for _, file := range filesInHomeDir {
		go func(f os.FileInfo) {
			time.Sleep(2 * time.Second)
			fmt.Printf("Visiting file %v\n", f.Name())
			defer wg.Done()
		}(file)
	}

	// We use Wait() to block until the WaitGroup counter becomes zero.
	fmt.Println("Waiting for goroutines...")
	wg.Wait()
}
