package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	loadEntireFile()
	fmt.Println()

	file := openFile()
	defer file.Close() // It’s idiomatic to defer a Close immediately after opening a file.

	readSomeBytes(file)
	seekLocation(file)
	ioRead(file)
	notBuiltInRewind(file)
	bufferedRead(file)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Perhaps the most basic file reading task is slurping a file's entire contents into memory.
func loadEntireFile() {
	data, err := ioutil.ReadFile("./file-read_sample.txt")
	check(err)
	fmt.Println(string(data))
}

// You'll often want more control over how and what parts of a file are read. For these tasks,
// start by `Open`ing a file to obtain an `os.File` value.
func openFile() *os.File {
	file, err := os.Open("./file-read_sample.txt")
	check(err)
	return file
}

// Read some bytes from the beginning of the file.
// Allow up to 5 to be read but also note how many actually were read.
func readSomeBytes(file *os.File) {
	bytes := make([]byte, 5)
	bytesRead, err := file.Read(bytes)
	check(err)
	fmt.Printf("%d bytes: %s\n", bytesRead, string(bytes))
}

// You can also `Seek` to a known location in the file and `Read` from there.
func seekLocation(file *os.File) {
	offset, err := file.Seek(6, 0)
	check(err)
	bytes := make([]byte, 2)
	bytesRead, err := file.Read(bytes)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", bytesRead, offset, string(bytes))
}

// The `io` package provides some functions that may be helpful for file reading.
// For example, reads like the ones above can be more robustly implemented with `ReadAtLeast`.
func ioRead(file *os.File) {
	offset, err := file.Seek(6, 0)
	check(err)
	bytes := make([]byte, 2)
	bytesRead, err := io.ReadAtLeast(file, bytes, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", bytesRead, offset, string(bytes))
}

// There is no built-in rewind, but `Seek(0, 0)` accomplishes this.
func notBuiltInRewind(file *os.File) {
	_, err := file.Seek(0, 0)
	check(err)
}

// The `bufio` package implements a buffered reader that may be useful both for its efficiency
// with many small reads and because of the additional reading methods it provides.
func bufferedRead(file *os.File) {
	bufioRead := bufio.NewReader(file)
	bytes, err := bufioRead.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(bytes))
}
