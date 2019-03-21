package main

import "fmt"

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	w1, w2 := multiple("Hey", "there")
	fmt.Println(w1, w2)
}
