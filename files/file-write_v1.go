package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	quickStart()

	file := moreGranular()
	defer file.Close() // It’s idiomatic to defer a Close immediately after opening a file.

	writeBytes(file)
	writeString(file)
	file.Sync() // Issue a Sync to flush writes to stable storage.
	bufferedIo(file)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// To start quickly, here’s how to dump a string (or just bytes) into a file.
func quickStart() {
	bytes := []byte("hello\ngo\n")
	var fileMode os.FileMode = 0644
	err := ioutil.WriteFile("./file-write_v1_sample-1.txt", bytes, fileMode)
	check(err)
	fmt.Printf("wrote %d bytes to file\n", len(bytes))
}

// For more granular writes, open a file for writing.
func moreGranular() *os.File {
	file, err := os.Create("./file-write_v1_sample-2.txt")
	check(err)
	return file
}

// You can Write byte slices as you’d expect.
func writeBytes(f *os.File) {
	bytes := []byte{115, 111, 109, 101, 10}
	bytesWritten, err := f.Write(bytes)
	check(err)
	fmt.Printf("wrote %d bytes to file\n", bytesWritten)
}

// A WriteString is also available.
func writeString(f *os.File) {
	bytesWritten, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes to file\n", bytesWritten)
}

// bufio provides buffered writers in addition to the buffered readers we saw earlier.
func bufferedIo(f *os.File) {
	bufioWriter := bufio.NewWriter(f)
	bytesWritten, err := bufioWriter.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes to file\n", bytesWritten)
	// Use Flush to ensure all buffered operations have been applied to the underlying writer.
	bufioWriter.Flush()
}
