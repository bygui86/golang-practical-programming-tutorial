package main

import "fmt"

func pointer(a *int) {
	fmt.Println("a *int:", a)
	fmt.Println("*a:", *a)
}

func value(c int) {
	fmt.Println("c int:", c)
	fmt.Println("&c:", &c)
}

func realVariable() {
	x := 5
	pointer(&x)
	// pointer(x) > cannot use x (type int) as type *int in argument to pointer
	fmt.Println("&x:", &x)
	value(x)
	// value(&x) > cannot use &x (type *int) as type int in argument to value
}

func referenceVariable() {
	x := 5
	a := &x
	// pointer(&a) > cannot use &a (type **int) as type *int in argument to pointer
	pointer(a)
	fmt.Println("&x:", &x)
	fmt.Println("&a:", &a)
	// value(a) > cannot use a (type *int) as type int in argument to value
	// value(&a) > cannot use &a (type **int) as type int in argument to value
	value(*a)
}

func main() {
	realVariable()
	fmt.Println()
	referenceVariable()
}
