package main

import "fmt"
import "math"

// return void
func printSqrt4() {
	fmt.Println("The square root of 4 is:", math.Sqrt(4))
}

// return value
func sqrt4() float64 {
	return math.Sqrt(4)
}

// return multiple values
func sqrt4msg() (string, float64) {
	return "The square root of 4 is:", math.Sqrt(4)
}

// return multiple named values
func namedSqrt4msg() (s string, f float64) {
	s = "The square root of 4 is:"
	f = math.Sqrt(4)
	return
}

func main() {
	printSqrt4()

	fmt.Println("The square root of 4 is:", sqrt4())

	fmt.Println(sqrt4msg())

	s, f := namedSqrt4msg()
	fmt.Println(s, f)
}
