package main

import "fmt"

func main() {
	fmt.Print("Print no newline")
	fmt.Println()
	fmt.Printf("%s", "Format print without newline")
	fmt.Println()
	fmt.Println("Print with newline")
	fmt.Println("First part", "-", "second part")
}
