package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("It is currently:", time.Now())
	fmt.Println()
	fmt.Println("A number from 1-100:", rand.Intn(100))
	fmt.Println()
	fmt.Println("Use 'go doc <PACKAGE> <METHOD>' to know more about a function")
}
