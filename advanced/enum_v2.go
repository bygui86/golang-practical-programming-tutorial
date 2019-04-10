package main

import (
	"fmt"
	"strconv"
)

// This is the base declaration of our enumeration
type Status int32

const (
	STARTING = Status(iota)
	// equivalent to the extended version
	// STARTING Status = iota
	RUNNING
	STOPPED
	LOCKED
	UNKNOWN
)

var (
	StatusNames = map[int32]string{
		0: "STARTING",
		1: "RUNNING",
		2: "STOPPED",
		3: "LOCKED",
		4: "UNKNOWN",
	}
	StatusValues = map[string]int32{
		"STARTING": 0,
		"RUNNING":  1,
		"STOPPED":  2,
		"LOCKED":   3,
		"UNKNOWN":  4,
	}
)

func (st Status) String() string {
	switch {
	case st <= UNKNOWN:
		return StatusNames[int32(st)]
	default:
		return strconv.Itoa(int(st))
	}
}

func main() {

	fromInt := Status(0)
	fmt.Println("Status FROM INT, string:", fromInt)
	fmt.Println("Status FROM INT, value:", int32(fromInt))
	fmt.Printf("Status FROM INT, type: %T", fromInt)
	fmt.Println()
	fmt.Println()

	fromConst := UNKNOWN
	fmt.Println("Status FROM CONST, string:", fromConst)
	fmt.Println("Status FROM CONST, value:", int32(fromConst))
	fmt.Printf("Status FROM CONST, type: %T", fromConst)
	fmt.Println()
	fmt.Println()

	fromString := Status(StatusValues["RUNNING"])
	fmt.Println("Status FROM STRING, string:", fromString)
	fmt.Println("Status FROM STRING, value:", int32(fromString))
	fmt.Printf("Status FROM STRING, type: %T", fromString)
}
