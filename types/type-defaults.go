package main

import "fmt"

/*
Go's basic types are:
	bool
	string
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // alias for uint8
	rune // alias for int32
		// represents a Unicode code point
	float32 float64
	complex64 complex128
*/

/*
The zero value is:
	0 for numeric types,
	false for the boolean type, and
	"" (the empty string) for strings.
*/

func main() {
	var str string
	var x float32
	var b bool
	var strArr []string
	var intArr []int
	strSlice := make([]string, 3)
	intSlice := make([]int, 5)

	fmt.Println("Type defaults:")
	fmt.Println("\tstring (empty):", str)
	fmt.Println("\tnumber:", x)
	fmt.Println("\tboolean:", b)
	fmt.Println("\tstring array:", strArr)
	fmt.Println("\tint array:", intArr)
	fmt.Println("\tstring slice:", strSlice) // empty strings
	fmt.Println("\tint slice:", intSlice)
}
