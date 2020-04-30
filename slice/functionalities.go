package main

import "fmt"

// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
func main() {
	s := createSlice()

	inlineDeclaration()

	getSet(s)

	lengthCapacity(s)

	appendTo(s)

	copyFrom(s)

	slicing(s)

	multiDimensional()
}

// To create an empty slice with non-zero length, use the builtin make.
// Here we make a slice of strings of length 3 (initially zero-valued).
func createSlice() []string {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	return s
}

// We can declare and initialize a variable for slice in a single line as well.
func inlineDeclaration() {
	t := []string{"g", "h", "i"}
	fmt.Println("inline declaration:", t)
}

// We can set and get just like with arrays.
func getSet(s []string) {
	fmt.Println("get before set:", s[2])
	s[0] = "a1"
	fmt.Println("get after set:", s[2])
}

func lengthCapacity(s []string) {
	// len returns the length of the slice
	// The length is the declared one creating the slice, it does not matter if the number
	// of objects inside it is the same.
	fmt.Println("len:", len(s))

	// cap returns the capacity of the slice (in this case same as length)
	fmt.Println("cap:", cap(s))
}

// In addition to these basic operations, slices support several more that
// make them richer than arrays. One is the builtin append, which returns a
// slice containing one or more new values. Note that we need to accept a return
// value from append as we may get a new slice value.
func appendTo(s []string) {
	fmt.Println("slice before append:", s)
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("slice after append:", s)
}

// Slices can also be copy’d. Here we create an empty slice c of the same
// length as s and copy into c from s.
func copyFrom(s []string) {
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("slice:", s)
	fmt.Println("copy:", c)
}

// Slices support a “slice” operator with the syntax slice[low:high].
// For example, this gets a slice of the elements s[2], s[3], and s[4].
func slicing(s []string) {
	l := s[2:5]
	fmt.Println("slice 1[2:5]:", l)

	// This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("slice 2[:5]:", l)

	// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("slice 3[2:]:", l)
}

// Slices can be composed into multi-dimensional data structures.
// The length of the inner slices can vary, unlike with multi-dimensional arrays.
func multiDimensional() {
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("multi dimensional:", twoD)
}
