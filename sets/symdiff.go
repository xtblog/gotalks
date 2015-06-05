package main

import (
	"fmt"
	"sort"

	"github.com/xtgo/set"
)

func main() {
	// BEGIN OMIT
	a, b := sort.IntSlice{2, 3, 4}, sort.IntSlice{1, 2, 4, 5}
	data := append(a, b...)

	size := set.SymDiff(data, len(a))
	data = data[:size]
	fmt.Println(data)
	// END OMIT
}
