package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {

	fmt.Println("Read files from folder:")
	readFiles(".", "file-collection")
}

func readFiles(strBasePath, strPath string) {

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

	for _, file := range files {
		fmt.Println("        ... reading file:", file.Name())

		fileContent, fileContentErr := ioutil.ReadFile(path + "/" + file.Name())
		if fileContentErr != nil {
			panic(fileContentErr)
		}

		fmt.Println(string(fileContent))
	}
}
