package main

import (
	"fmt"
)

func main() {
	integer := 3
	fmt.Println("int:", integer)
	fmt.Println("int32:", int32(integer))
	fmt.Println("float64:", float64(integer))

	floating := 3.14159
	fmt.Println("float64:", floating)
	fmt.Println("int32:", int32(floating))
	fmt.Println("int:", int(integer))
}
