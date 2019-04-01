package main

import (
	// Go implements several hash functions in various crypto/* packages.
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

/*
	SHA1 hashes are frequently used to compute short identities for binary or text blobs.
	For example, the git revision control system uses SHA1s extensively to identify
	versioned files and directories.
*/
func main() {
	str := "hash this string"

	// The pattern for generating a hash is ALGO.New(), ALGO.Write(bytes), then ALGO.Sum([]byte{}).
	computeSha1(&str)
	fmt.Println()
	computeMd5(&str)
}

func computeSha1(str *string) {

	// Here we start with a new hash.
	hash := sha1.New()

	// Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.
	bytes := []byte(*str)
	hash.Write(bytes)

	// This gets the finalized hash result as a byte slice. The argument to Sum can be used to
	// append to an existing byte slice: it usually isn’t needed.
	bytesSlice := hash.Sum(nil)

	// SHA1 values are often printed in hex, for example in git commits. Use the %x format
	// verb to convert a hash results to a hex string.
	fmt.Println(*str)
	fmt.Printf("SHA1: %x\n", bytesSlice)
}

func computeMd5(str *string) {

	// Here we start with a new hash.
	hash := md5.New()

	// Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.
	bytes := []byte(*str)
	hash.Write(bytes)

	// This gets the finalized hash result as a byte slice. The argument to Sum can be used to
	// append to an existing byte slice: it usually isn’t needed.
	bytesSlice := hash.Sum(nil)

	// SHA1 values are often printed in hex, for example in git commits. Use the %x format
	// verb to convert a hash results to a hex string.
	fmt.Println(*str)
	fmt.Printf("MD5: %x\n", bytesSlice)
}
