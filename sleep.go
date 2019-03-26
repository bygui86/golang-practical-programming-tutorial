package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var msg string = "Message number "
	var counter int = 1
	for {
		fmt.Println(msg + strconv.Itoa(counter))
		counter++
		time.Sleep(3 * time.Second)
	}
}
