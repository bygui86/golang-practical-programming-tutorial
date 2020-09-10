package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

/*
	HOW TO RUN

	1. run benchmark
		go test -bench . -count 10 > buffered-bench-results.txt

	2. analyse results with benchstat
		benchstat buffered-bench-results.txt


	LINKS

		https://medium.com/a-journey-with-go/go-buffered-and-unbuffered-channels-29a107c00268
		https://blog.jiayu.co/2019/05/benchmarking-go-code/
		https://pkg.go.dev/golang.org/x/perf/cmd/benchstat?tab=doc
*/

func BenchmarkWithNoBuffer(b *testing.B) {
	benchmarkWithBuffer(b, 0)
}

func BenchmarkWithBuffer_SizeOf1(b *testing.B) {
	benchmarkWithBuffer(b, 1)
}

func BenchmarkWithBuffer_SizeEqualsToNumberOfWorker(b *testing.B) {
	benchmarkWithBuffer(b, 5)
}

func BenchmarkWithBuffer_SizeExceedsNumberOfWorker(b *testing.B) {
	benchmarkWithBuffer(b, 25)
}

func benchmarkWithBuffer(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		c := make(chan uint32, size)

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()

			for i := uint32(0); i < 1000; i++ {
				c <- i % 2
			}
			close(c)
		}()

		var total uint32
		for w := 0; w < 5; w++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				for {
					v, ok := <-c
					if !ok {
						break
					}
					atomic.AddUint32(&total, v)
				}
			}()
		}

		wg.Wait()
	}
}
