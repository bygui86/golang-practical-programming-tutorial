package main

import "fmt"

// var VAR_NAME map[KEY_TYPE]VALUE_TYPE
// In Go, maps are REFERENCE TYPES, and they point to anything with value.
// You cannot write to a map not initialized, because it points nowhere. Before we have to intitialize it with 'make'
func main() {
	grades := make(map[string]float32)
	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 71
	fmt.Println("grades map:")
	for k, v := range grades {
		fmt.Println("\t", k, ":", v)
	}

	tims_grade := grades["Timmy"]
	fmt.Println("Timmy:", tims_grade)

	fmt.Println("Removed Timmy from grades map")
	delete(grades, "Timmy")
	fmt.Println(grades)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	value, ok := kvs["c"]
	if ok {
		fmt.Println("Value PRESENT:", value)
	} else {
		fmt.Println("Value NOT PRESENT")
	}
}
