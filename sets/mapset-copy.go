package main

import (
	"fmt"
	"sort"
)

// BEGIN OMIT
type StringSet map[string]bool

// Intersect returns a copy containing the intersection of both sets.
func (s StringSet) Intersect(t StringSet) StringSet {
	m := make(StringSet)
	for k := range s {
		if t[k] {
			m[k] = true
		}
	}
	return m
}

func main() {
	x := StringSet{"a": true, "b": true}
	y := StringSet{"b": true, "c": true}
	z := x.Intersect(y)
	fmt.Printf("%v & %v == %v\n", x, y, z)
}

// END OMIT

func (s StringSet) String() string {
	lst := make(sort.StringSlice, 0, len(s))
	for k, ok := range s {
		if ok {
			lst = append(lst, k)
		}
	}
	sort.Sort(lst)
	return fmt.Sprint(lst)
}
