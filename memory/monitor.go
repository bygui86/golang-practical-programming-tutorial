package main

import (
	"fmt"
	"runtime"
	"time"
)

type Monitor struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects,
	PauseTotalNs uint64
	NumGC        uint32
	NumGoroutine int
}

// See https://scene-si.org/2018/08/06/basic-monitoring-of-go-apps-with-the-runtime-package/
func main() {
	duration := 1
	var monitor Monitor
	var memStats runtime.MemStats
	for {
		// every second
		<-time.After(time.Duration(duration) * time.Second)

		// Read full mem stats
		runtime.ReadMemStats(&memStats)

		// Number of goroutines
		monitor.NumGoroutine = runtime.NumGoroutine()

		// Misc memory stats
		monitor.Alloc = memStats.Alloc
		monitor.TotalAlloc = memStats.TotalAlloc
		monitor.Sys = memStats.Sys
		monitor.Mallocs = memStats.Mallocs
		monitor.Frees = memStats.Frees

		// Live objects = Mallocs - Frees
		monitor.LiveObjects = monitor.Mallocs - monitor.Frees

		// GC Stats
		monitor.PauseTotalNs = memStats.PauseTotalNs
		monitor.NumGC = memStats.NumGC

		// print out ...
		fmt.Printf("Monitor: %+v\n", monitor)

		// ... or encode to json and print/send/whatever ...
		// b, _ := json.Marshal(monitor)
		// fmt.Printf("Monitor json: %s\n", string(b))
	}
}
