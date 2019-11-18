package main

import (
	"fmt"
	"sync"
)

/*
	It allows you to define a task which you only want to execute once during the lifetime of your program.
	This is very useful for Singleton-like behavior.
	It has a single Do function that lets you pass another function which you intend to execute only once.

	When you first access it, it will be a little slow in returning and you will see the following logs in the server:
		http handler start
		one time op start
		one time op end
		http handler end
*/

var once sync.Once

func main() {
	DoSomething()
	DoSomething()
}

func DoSomething() {
	once.Do(func() {
		fmt.Println("Run once - first time, loading...")
	})
	fmt.Println("Run this every time")
}
