package main

import "fmt"

func add_extended(x float64, y float64) float64 {
	return x + y
}
func add_short(x, y float64) float64 {
	return x + y
}

func sum_extended() {
	var num1 float64 = 5.6
	var num2 float64 = 9.5
	fmt.Println(add_short(num1, num2))
}
func sum_short() {
	var num1, num2 float64 = 5.6, 9.5
	fmt.Println(add_short(num1, num2))
}
func sum_evenshorter() {
	num1, num2 := 5.6, 9.5
	fmt.Println(add_short(num1, num2))
}

func main() {
	sum_extended()
	sum_short()
	sum_evenshorter()
}
