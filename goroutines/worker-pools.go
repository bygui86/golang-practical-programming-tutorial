package main

import (
	"fmt"
	"runtime"
	"time"
)

// Here we implement a worker pool using goroutines and channels.
func main() {
	numCPUs := runtime.NumCPU()

	// In order to use our pool of workers we need to send them work and collect their results.
	// We make 2 channels for this.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// This starts up as many workers as CPUs of your machine, initially blocked because
	// there are no jobs yet.
	for w := 1; w <= numCPUs; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 24 jobs and then close that channel to indicate that’s all the work we have.
	for j := 1; j <= 24; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work.
	for r := 1; r <= 24; r++ {
		<-results
	}
}

/*
Here’s the worker, of which we’ll run several concurrent instances.
These workers will receive work on the jobs channel and send the corresponding results on results.
We’ll sleep a second per job to simulate an expensive task.
*/
func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
