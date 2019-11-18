package main

import (
	"fmt"
	"strconv"
)

// Here we simulate C’s enum by Go’s iota
type Season uint8

const (
	Spring = Season(iota)
	Summer
	Autumn
	Winter
)

// output function for Season variable
func (s Season) String() string {
	name := []string{"spring", "summer", "autumn", "winter"}
	i := uint8(s)
	switch {
	case i <= uint8(Winter):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}

func main() {

	validConst := Summer
	fmt.Println("Valid const:", validConst)

	validInt := Season(1)
	fmt.Println("Valid int enum:", validInt)

	invalid := Season(9)
	fmt.Println("Invalid int enum:", invalid)
}
