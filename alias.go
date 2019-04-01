package main

import (
	"fmt"
	s "strings" // Here we can define an alias for a package
)

// Here we can define an alias for a function
var p = fmt.Println

func main() {
	// We use 's' instead of 'strings'
	// We use 'p' instaed of 'fmt.Println'
	p("Contains:  ", s.Contains("test", "es"))
	p()
}
