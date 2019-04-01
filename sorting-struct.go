package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type byAge []person

func (s byAge) Len() int {
	return len(s)
}
func (s byAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byAge) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

// Go’s sort package implements sorting for builtins and user-defined types.
func main() {
	matt := person{"Matt", 33}
	john := person{"John", 42}
	jane := person{"Jane", 29}

	persons := []person{matt, jane, john}
	usingFunction(persons)
	fmt.Println()
	usingPointerReceiver(persons)
}

func usingFunction(persons []person) {
	fmt.Println("Persons:", persons)
	// Use 'persons[:]' for an array and 'persons' for a slice
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Age < persons[j].Age
	})
	fmt.Println("Persons sorted by age:", persons)
}

func usingPointerReceiver(persons []person) {
	fmt.Println("Persons:", persons)
	sort.Sort(byAge(persons))
	fmt.Println("Persons sorted by age:", persons)
}
