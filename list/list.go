package main

import (
	"container/list"
	"fmt"
)

type sample struct {
	x int
}

func main() {
	// New list
	values := list.New()

	// Add 3 elements at the front
	values.PushFront("bird")
	values.PushFront(sample{5})
	values.PushFront(true)

	// Add 5 elements at the back
	for i := 0; i < 5; i++ {
		values.PushBack(i)
	}

	// Features
	fmt.Println("Length:", values.Len())
	fmt.Println("Front:", values.Front().Value)
	fmt.Println("Back:", values.Back().Value)

	// Loop over
	for temp := values.Front(); temp != nil; temp = temp.Next() {
		fmt.Println(temp.Value)
	}
}
