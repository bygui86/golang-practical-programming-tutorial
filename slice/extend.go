package main

import "fmt"

func main() {
	extend()
	fmt.Println()
	extendAndCopy()
}

func extend() {
	a := createSlice(5)
	fmt.Println("a", len(a), cap(a), a)

	extension := 10
	a = append(a, make([]int, extension)...)

	fmt.Println("extended a", len(a), cap(a), a)
}

func extendAndCopy() {
	a := createSlice(5)
	b := createSlice(5)
	fmt.Println("a", len(a), cap(a), a)
	fmt.Println("b", len(b), cap(b), b)

	a = append(a, b...)

	fmt.Println("extended a", len(a), cap(a), a)
	fmt.Println("b", len(b), cap(b), b)
}

func createSlice(length int) []int {
	s := make([]int, length)
	for i := 0; i < length; i++ {
		s[i] = i
	}
	return s
}
