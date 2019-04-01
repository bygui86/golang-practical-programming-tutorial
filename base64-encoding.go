package main

import (
	"encoding/base64"
	"fmt"
)

// Go provides built-in support for base64 encoding/decoding.
// Go supports both standard and URL-compatible base64.
func main() {
	data := "abc123!?$*&()'-=@~"
	// The encoder requires a []byte so we cast our string to that type.
	bytes := []byte(data)

	fmt.Println("Source string:", data)
	fmt.Println()
	standardEncoding(&bytes)
	fmt.Println()
	urlEncoding(&bytes)
}

// Here’s how to encode/decode using the standard encoder.
func standardEncoding(bytes *[]byte) {
	stdEncoded := base64.StdEncoding.EncodeToString(*bytes)
	fmt.Println("Standard encoding:", stdEncoded)

	stdDecoded, _ := base64.StdEncoding.DecodeString(stdEncoded)
	fmt.Println("Standard decoding:", string(stdDecoded))
}

// Here’s how to encodes/decodes using a URL-compatible base64 format.
func urlEncoding(bytes *[]byte) {

	urlEncoded := base64.URLEncoding.EncodeToString(*bytes)
	fmt.Println("URL encoding:", urlEncoded)

	urlDecoded, _ := base64.URLEncoding.DecodeString(urlEncoded)
	fmt.Println("URL decoding:", string(urlDecoded))
}
