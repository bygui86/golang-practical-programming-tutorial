package main

import "fmt"

// different than type IntSlice []int
type IntSlice = []int

func main() {
	// intSlice := []int{1, 2}
	intSlice := IntSlice{1, 2}
	fmt.Printf("slice pointer: %p\n", &intSlice)
	passAliasByReference(&intSlice)
	passSliceByReference(&intSlice)
	fmt.Println()

	intArray := [...]int{1, 2}
	fmt.Printf("array pointer: %p\n", &intArray)
	passByCopy(intArray)
}

func passAliasByReference(slice *IntSlice) {
	fmt.Printf("pointer: %p\n", slice)
	fmt.Println("value:", *slice)
}

func passSliceByReference(slice *[]int) {
	fmt.Printf("pointer: %p\n", slice)
	fmt.Println("value:", *slice)
}

func passByCopy(array [2]int) {
	fmt.Printf("pointer: %p\n", &array)
	fmt.Println("value:", array)
}
