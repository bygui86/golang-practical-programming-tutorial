package main

import (
	"fmt"
	"time"
)

const (
	// See http://golang.org/pkg/time/#Parse
	timeFormat = "2006-01-02 15:04 MST"
)

func main() {
	// parse
	v := "2014-05-03 20:57 UTC"
	then, err := time.Parse(timeFormat, v)
	if err != nil {
		fmt.Println(err)
		return
	}

	// since
	duration := time.Since(then)
	fmt.Println(duration.Hours())

	// difference
	t1 := time.Now()
	t2 := t1.Add(time.Second * 341)
	diff := t2.Sub(t1)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(diff)
}
