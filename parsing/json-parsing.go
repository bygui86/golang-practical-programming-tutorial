package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// We’ll use these two structs to demonstrate encoding and decoding of custom types below.
type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"p"`
	Fruits []string `json:"f"`
}

/*
	Go offers built-in support for JSON encoding and decoding,
	including to and from built-in and custom data types.
*/
func main() {
	// ENCODING
	encodeBasicTypes()
	fmt.Println()
	encodeCustomTypes()
	fmt.Println()

	// DECODING
	decodeGenericStructure()
	fmt.Println()
	decodeToStruct()
	fmt.Println()

	// STREAM
	encodeToStream()
}

func encodeBasicTypes() {
	// Atomic values
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// Slices and maps, which encode to JSON arrays and objects as expect
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
}

func encodeCustomTypes() {
	/*
		The JSON package can automatically encode your custom data types.
		It will only include exported fields in the encoded output and will by
		default use those names as the JSON keys.
	*/
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// You can use tags on struct field declarations to customize the encoded JSON key names.
	// Check the definition of response2 above to see an example of such tags.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
}

func decodeGenericStructure() {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// We need to provide a variable where the JSON package can put the decoded data.
	// This map[string]interface{} will hold a map of strings to arbitrary data types.
	var dat map[string]interface{}
	// Here’s the actual decoding, and a check for associated errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// In order to use the values in the decoded map, we’ll need to cast them to their
	// appropriate type. For example here we cast the value in num to the expected float64 type.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Accessing nested data requires a series of casts.
	strs := dat["strs"].([]interface{})
	strs0 := strs[0].(string)
	fmt.Println(strs0)
}

func decodeToStruct() {
	/*
		We can also decode JSON into custom data types. This has the advantages of adding
		additional type-safety to our programs and eliminating the need for type assertions
		when accessing the decoded data.
	*/
	str := `{"p": 1, "f": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}

func encodeToStream() {
	// We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
