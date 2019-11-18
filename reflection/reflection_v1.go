package main

import (
	"fmt"
	"reflect"
	"time"
)

type custom struct {
	Name string
	Age  int
}

func main() {

	strV := "string"
	intV := 10
	floatV := 1.2
	timeV := time.Now()
	customV := custom{"matt", 33}

	fmt.Println(reflect.TypeOf(strV))
	fmt.Println(reflect.TypeOf(intV))
	fmt.Println(reflect.TypeOf(floatV))
	fmt.Println(reflect.TypeOf(timeV))
	fmt.Println(reflect.TypeOf(customV))
}
