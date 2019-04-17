package main

import (
	"fmt"
	"time"
)

/*
	We often want to execute Go code at some point in the future, or repeatedly at some interval.
	Go’s built-in timer and ticker features make both of these tasks easy.

	Timers represent a single event in the future.
	You tell the timer how long you want to wait, and it provides a channel that will be notified
	at that time.
*/
func main() {
	fmt.Println("Timer samples...")

	// This timer will wait 3 seconds.
	timer1 := time.NewTimer(3 * time.Second)

	// The <-timer1.C blocks on the timer’s channel C until it sends a value indicating
	// that the timer expired.
	fmt.Print("Main waiting for Timer 1 to be expired ... ")
	<-timer1.C
	fmt.Println("EXPIRED")

	// But if you just wanted to wait, you could have used time.Sleep.
	// One reason a timer may be useful is that you can cancel the timer before it expires.
	// Here’s an example of that.
	timer2 := time.NewTimer(time.Second)
	go func() {
		fmt.Print("GoRoutine waiting for Timer 2 to be expired ... ")
		<-timer2.C
		fmt.Println("EXPIRED")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("STOPPED")
	}

	/* PLEASE NOTE
	The first timer will expire ~3s after we start the program.
	The second should be stopped before it has a chance to expire.
	*/
}
