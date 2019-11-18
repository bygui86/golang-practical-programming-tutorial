package main

import (
	"os"
)

/*
	A common use of panic is to abort if a function returns an error value that we don’t know
	how to (or want to) handle. Here’s an example of panicking if we get an unexpected error
	when creating a new file.

	Running this program will cause it to panic, print an error message and goroutine traces,
	and exit with a non-zero status.
*/
func main() {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
