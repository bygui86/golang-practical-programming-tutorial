package main

import "fmt"

// different than type IntSlice []int
type IntSlice = []int

func main() {
	intSlice := IntSlice{1, 2}
	fmt.Printf("slice: %+v\n", intSlice)
	fmt.Printf("slice pointer: %p\n", &intSlice)

	passByReference(&intSlice)
	fmt.Println()

	passByCopy(intSlice)
	fmt.Println()

	passByAliasRef(&intSlice)
	fmt.Println()

	passByAliasCopy(intSlice)
	fmt.Println()
}

func passByReference(slice *[]int) {
	fmt.Println("slice:", *slice)
	fmt.Printf("pointer: %p\n", slice)
}

func passByCopy(slice []int) {
	fmt.Println("slice:", slice)
	fmt.Printf("pointer: %p\n", &slice)
}

func passByAliasRef(slice *IntSlice) {
	fmt.Println("slice:", *slice)
	fmt.Printf("pointer: %p\n", slice)
}

func passByAliasCopy(slice IntSlice) {
	fmt.Println("slice:", slice)
	fmt.Printf("pointer: %p\n", slice)
}
