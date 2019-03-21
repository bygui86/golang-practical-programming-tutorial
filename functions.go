package main

import "fmt"
import "math"

// return void
func print_sqrt4() {
	fmt.Println("The square root of 4 is:", math.Sqrt(4))
}

// return value
func sqrt4() float64 {
	return math.Sqrt(4)
}

func main() {
	print_sqrt4()
	fmt.Println("The square root of 4 is:", sqrt4())
}
