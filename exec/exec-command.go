package main

import (
	"fmt"
	"os/exec"
)

func main() {

	out, err := exec.Command("date", "+%Y-%m-%d").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Printf("Current date is %s", out)
	fmt.Println()
}
