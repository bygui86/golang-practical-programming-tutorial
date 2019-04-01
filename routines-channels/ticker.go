package main

import (
	"fmt"
	"time"
)

/*
	We often want to execute Go code at some point in the future, or repeatedly at some interval.
	Go’s built-in timer and ticker features make both of these tasks easy.
	Tickers are for when you want to do something repeatedly at regular intervals.
*/
func main() {
	// Tickers use a similar mechanism to timers: a channel that is sent values.
	// Here we’ll use the range builtin on the channel to iterate over the values as
	// they arrive every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Tickers can be stopped like timers. Once a ticker is stopped it won’t receive
	// any more values on its channel. We’ll stop ours after 2100ms.
	time.Sleep(2100 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
