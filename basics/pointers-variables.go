package main

import "fmt"

func main() {
	x := 15
	fmt.Println("Value of x:", x) // prints out the mem addr

	a := &x                           // point to x (memory address)
	fmt.Println("Reference of a:", a) // prints out the mem addr
	fmt.Println("Value of a:", *a)    // read a through the pointer, so this will print out a value (15 in this case)

	*a = 5                                // sets the value pointed at to 5, which means x is modified (since x is stored at the mem addr)
	fmt.Println("Second value of a:", *a) // see the new value of a
	fmt.Println("Second value of x:", x)  // see the new value of x

	*a = *a * *a                        // what is the value of x now?
	fmt.Println("Third value of x", x)  // prints out the value
	fmt.Println("Third value of a", *a) // prints out the value

	fmt.Println("Reference of x:", &x) // prints out the mem addr
	fmt.Println("Reference of a:", a)  // prints out the mem addr
}
