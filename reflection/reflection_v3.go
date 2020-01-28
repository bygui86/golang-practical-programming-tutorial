package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	completeSample()
}

func completeSample() {
	power := func(x int) (int, error) {
		if x < 0 {
			return 0, errors.New("x should not be negative")
		}
		return x * x, nil
	}

	t := reflect.TypeOf(power) // t.Kind() == reflect.Func

	fmt.Println("Kind:", t.Kind())             // kind: func
	fmt.Println("Num of inputs:", t.NumIn())   // num inputs: 1
	fmt.Println("Num of outputs:", t.NumOut()) // num outputs: 2
	fmt.Println("Type of input 0:", t.In(0))   // input types: int
	fmt.Println("Type of output 0:", t.Out(0)) // output types: int, error
	fmt.Println("Type of output 1:", t.Out(1))

	in := make([]reflect.Value, 0)
	in = append(in, reflect.ValueOf(5))

	v := reflect.ValueOf(power)
	results := v.Call(in) // dynamic dispatch
	out := results[0].Int()
	var err error
	if !results[1].IsNil() {
		err = results[1].Interface().(error)
	}
	fmt.Printf("Output for input 5: %d, %v", out, err)
	fmt.Println("")
}
