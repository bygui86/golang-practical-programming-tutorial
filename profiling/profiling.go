package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/pkg/profile"
)

// see https://flaviocopes.com/golang-profiling/
/*
	INSTRUCTIONS:
		. set go-modules
			export GO111MODULE=on
		. build
			go build
		. run
			go run profiling.go
		. stop
			ctrl + c
		. install graphviz (https://www.graphviz.org)
		. generate the analysis using 'go tool pprof'
			go tool pprof --{format} -source_path . *.pprof > {file_name}.{format}
			e.g.
				go tool pprof --pdf -source_path . mem.pprof > mem.pdf
				go tool pprof --txt -source_path . mem.pprof > mem.txt
		. open {file_name}.{format}
		. analyse!
*/
func main() {
	// CPU profiling by default
	defer profile.Start(profile.ProfilePath(".")).Stop()

	// WARN: you can run just ONE PROFILER at a time

	// Available profilers
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.BlockProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.GoroutineProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MutexProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.ThreadcreationProfile, profile.ProfilePath(".")).Stop()

	doStuffs()
}

func doStuffs() {
	PrintMemUsage()

	var overall [][]int
	for i := 0; i < 4; i++ {
		// Allocate memory using make() and append to overall (so it doesn't get
		// garbage collected). This is to create an ever increasing memory usage
		// which we can track. We're just using []int as an example.
		a := make([]int, 0, 999999)
		overall = append(overall, a)

		// Print our memory usage at each interval
		PrintMemUsage()
		time.Sleep(time.Second)
	}

	// Clear our memory and print usage, unless the GC has run 'Alloc' will remain the same
	overall = nil
	PrintMemUsage()

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
