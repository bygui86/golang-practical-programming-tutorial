package main

import (
	"fmt"
	"time"
)

func intSwitch() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}

func timeSwitch() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}

func noConditionSwitch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}

func stringSwitch(s string) {
	switch s {
	case "one":
		fmt.Println("I'm number 1")
	case "two":
		fmt.Println("I'm number 2")
	case "three":
		fmt.Println("I'm number 3")
	default:
		fmt.Println("I'm another number")

	}
}

func typeSwitch(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("I'm a bool")
	case int:
		fmt.Println("I'm an int")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}
}

func main() {
	intSwitch()
	timeSwitch()
	noConditionSwitch()
	stringSwitch("two")
	typeSwitch(1)
}
