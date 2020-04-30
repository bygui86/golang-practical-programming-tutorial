package main

import "fmt"

func main() {
	expand()
	fmt.Println()
	expandAndCopy()
}

func expand() {
	a := createSlice(5)
	fmt.Println("a", len(a), cap(a), a)

	appendPosition := len(a)
	expansion := 10
	a = append(a[:appendPosition], append(make([]int, expansion), a[appendPosition:]...)...)

	fmt.Println("expanded a", len(a), cap(a), a)
}

func expandAndCopy() {
	a := createSlice(5)
	b := createSlice(5)
	fmt.Println("a", len(a), cap(a), a)
	fmt.Println("b", len(b), cap(b), b)

	appendPosition := 5
	a = append(a[:appendPosition], append(b, a[appendPosition:]...)...)

	fmt.Println("expanded a", len(a), cap(a), a)
	fmt.Println("b", len(b), cap(b), b)
}

func createSlice(length int) []int {
	s := make([]int, length)
	for i := 0; i < length; i++ {
		s[i] = i
	}
	return s
}
