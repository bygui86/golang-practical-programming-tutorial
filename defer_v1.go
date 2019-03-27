package main

import (
	"fmt"
)

// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
func main() {

	defer fmt.Println("this message will be printed as soon as the main ends... even if it comes other statements")

	fmt.Println("this message will be printed regurarly")
}
