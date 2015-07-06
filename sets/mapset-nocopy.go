package main

import (
	"fmt"
	"sort"
)

type StringSet map[string]bool

// BEGIN OMIT
// Intersect updates s in-place to contain the intersection of both sets.
func (s StringSet) IntersectInPlace(t StringSet) {
	for k := range s {
		if !t[k] {
			delete(s, k)
		}
	}
}

func (s StringSet) Copy() StringSet {
	m := make(StringSet, len(s))
	for k := range s {
		m[k] = true // HL
	}
	return m
}

func main() {
	x := StringSet{"a": true, "b": true} // try with {"c": false} // HL
	y := StringSet{"b": true, "c": true}
	z := x.Copy()
	z.IntersectInPlace(y)
	fmt.Printf("%v & %v == %v\n", x, y, z)
}

// END OMIT

func (s StringSet) String() string {
	lst := make(sort.StringSlice, 0, len(s))
	for k := range s {
		lst = append(lst, k)
	}
	sort.Sort(lst)
	return fmt.Sprint(lst)
}
