package main

import (
	"fmt"
	"sort"

	"github.com/xtgo/set"
)

func main() {
	// BEGIN OMIT
	a := sort.IntSlice{2, 3, 5, 7, 11, 13}   // primes
	b := sort.IntSlice{0, 1, 2, 3, 5, 8, 13} // Fibonacci numbers
	c := sort.IntSlice{1, 2, 5, 14, 42}      // Catalan numbers
	data := append(append(a, b...), c...)
	pivots := set.Pivots(len(a), len(b), len(c))

	size := set.Apply(set.Inter, data, pivots)
	data = data[:size]
	fmt.Println(data)
	// END OMIT
}
