package main

import "github.com/bygui86/golang-practical-programming-tutorial/goroutines/gopher-farm/farm"

// see https://medium.com/@olena_stoliarova/introduction-to-concurrency-in-go-gopher-farm-3aa23f253420
func main() {
	farm := farm.Build()

	farm.Open()

	farm.LiveFreeOrDieHard()
}
