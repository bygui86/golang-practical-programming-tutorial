package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {

	fmt.Println("Absolute path:")
	absolutePath("/Users/matteobaiguini/projects")

	fmt.Println("Absolute path with a relative path (results should be same as next text with a relative path):")
	absolutePath("file-collection")

	fmt.Println("Relative path:")
	relativePath(".", "file-collection")
}

// WARN: if the filepath.Abs(..) is given a relative path as input, it will consider the folder of the go file as base path
func absolutePath(strPath string) {

	path, pathErr := filepath.Abs(strPath)
	if pathErr != nil {
		panic(pathErr)
	}
	fmt.Println("    ... path:", path)

	files, filesErr := ioutil.ReadDir(path)
	if filesErr != nil {
		panic(filesErr)
	}
	fmt.Println("    ...", len(files), "files found")
}

func relativePath(strBasePath, strPath string) {

	path, pathErr := filepath.Rel(strBasePath, strPath)
	if pathErr != nil {
		panic(pathErr)
	}
	fmt.Println("    ... path:", path)

	files, filesErr := ioutil.ReadDir(path)
	if filesErr != nil {
		panic(filesErr)
	}
	fmt.Println("    ...", len(files), "files found")
}
