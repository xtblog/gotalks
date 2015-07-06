package main

import "fmt"

func main() {
	// BEGIN OMIT
	set := make(map[string]bool)
	set["a"] = true
	set["b"] = true

	// iterate over set elements (unordered)
	for elem := range set {
		fmt.Println(elem)
	}

	fmt.Println("c in set?", set["c"])
	// END OMIT
}
