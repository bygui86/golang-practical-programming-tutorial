package main

import "fmt"

// ... represents a variadic
func Sum(nums ...int) int {

	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

// see https://yourbasic.org/golang/three-dots-ellipsis/
func main() {

	// In an array literal, the ... notation specifies a length equal to the number of elements in the literal.
	stooges := [...]string{"Moe", "Larry", "Curly"}
	fmt.Println(len(stooges)) // 3

	// ...
	primes := []int{2, 3, 5, 7}
	// ... converts an array into a variadic
	fmt.Println(Sum(primes...)) // 17
}
