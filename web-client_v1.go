package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// With the Go programming language, we can't define variables that we wont use, which can be unfortunate with function returns.
	// resp, err := http.Get("https://google.com")
	// In order to ignore the possible error return, we can use underscore ('_')
	resp, _ := http.Get("https://google.com")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println("Google homepage response body:", string_body)
	// Don't forget to close to free up resources!
	resp.Body.Close()
}
