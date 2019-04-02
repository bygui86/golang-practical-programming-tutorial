package main

import (
	"bytes"
	"fmt"
	"regexp"
)

type sample struct {
	x int
}

// Go offers built-in support for regular expressions.
func main() {
	// This tests whether a pattern matches a string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Above we used a string pattern directly, but for other regexp tasks you’ll need to 'compile'
	// an optimized Regexp struct.
	// Many methods are available on these structs.
	compiledRegex, _ := regexp.Compile("p([a-z]+)ch")

	// Here’s a match test like we saw earlier.
	fmt.Println(compiledRegex.MatchString("peach"))

	// This finds the match for the regexp.
	fmt.Println(compiledRegex.FindString("peach punch"))

	// This also finds the first match but returns the start and end indexes for the match instead
	// of the matching text.
	fmt.Println(compiledRegex.FindStringIndex("peach punch"))

	// The Submatch variants include information about both the whole-pattern matches and the
	// submatches within those matches. For example this will return information for both
	// p([a-z]+)ch and ([a-z]+).
	fmt.Println(compiledRegex.FindStringSubmatch("peach punch"))

	// Similarly this will return information about the indexes of matches and submatches.
	fmt.Println(compiledRegex.FindStringSubmatchIndex("peach punch"))

	// The All variants of these functions apply to all matches in the input, not just the first.
	// For example to find all matches for a regexp.
	fmt.Println(compiledRegex.FindAllString("peach punch pinch", -1))

	// These All variants are available for the other functions we saw above as well.
	fmt.Println(compiledRegex.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// Providing a non-negative integer as the second argument to these functions will limit the
	// number of matches.
	fmt.Println(compiledRegex.FindAllString("peach punch pinch", 2))

	// Our examples above had string arguments and used names like MatchString. We can also
	// provide []byte arguments and drop String from the function name.
	fmt.Println(compiledRegex.Match([]byte("peach")))

	// When creating constants with regular expressions you can use the MustCompile variation
	// of Compile. A plain Compile won’t work for constants because it has 2 return values.
	constCompiledRegex := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(constCompiledRegex)

	// The regexp package can also be used to replace subsets of strings with other values.
	fmt.Println(constCompiledRegex.ReplaceAllString("a peach", "<fruit>"))

	// The Func variant allows you to transform matched text with a given function.
	in := []byte("a peach")
	out := constCompiledRegex.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
