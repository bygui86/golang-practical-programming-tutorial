package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("./file-write_v2_sample.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("create file")

	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("write on file")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("close file")
	f.Close()
}
