package main

import (
	"flag"
	"fmt"
)

func main() {
	// Basic flag declarations are available for string, integer, and boolean options.
	// Here we declare a string flag word with a default value "foo" and a short description.
	// WARNING: This flag.String function returns a string pointer, not a string value!
	word := flag.String("word", "", "a string [REQUIRED]")

	// This declares numb and fork flags, using a similar approach to the word flag.
	numb := flag.Int("numb", 42, "an int [REQUIRED]")
	boolean := flag.Bool("boolean", false, "a boolean [OPTIONAL]")

	// It’s also possible to declare an option that uses an existing var declared
	// elsewhere in the program. Note that we need to pass in a pointer to the flag declaration function.
	var svar string
	flag.StringVar(&svar, "refword", "", "a string by reference [OPTIONAL]")

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	if len(*word) == 0 {
		panic("no word defined, please set the -word flag")
	}

	if *numb == -1 {
		panic("no numb defined, please set the -numb flag")
	}

	fmt.Println("word:", *word)
	fmt.Println("numb:", *numb)
	fmt.Println("boolean:", *boolean)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
