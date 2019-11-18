package main

import (
	"fmt"
	"reflect"
)

func main() {
	typeOfSample()
	// interfaceToReflect()
	modificationProhibited()
	modifyThroughReflection()
}

func typeOfSample() {
	fmt.Printf("%T", 3) //int
}

func interfaceToReflect() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	v.SetFloat(7.1) // Error: will panic
}

func modificationProhibited() {
	var x float64 = 3.4
	y := reflect.ValueOf(&x)
	fmt.Println("type of y", y.Type())           // *float64
	fmt.Println("settability of y:", y.CanSet()) // false
}

func modifyThroughReflection() {
	var x float64 = 3.4
	y := reflect.ValueOf(&x)
	fmt.Println("type of y", y.Type())           // *float64
	fmt.Println("settability of y:", y.CanSet()) // false
	z := y.Elem()
	z.SetFloat(7.1)
	fmt.Println(z.Interface()) // 7.1
	fmt.Println(x)             // 7.1
}
