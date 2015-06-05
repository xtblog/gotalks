package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// BEGIN OMIT
func Shuffle(data sort.Interface) { sort.Sort(proxy{data, rand.Perm(data.Len())}) }

type proxy struct {
	data sort.Interface
	sort.IntSlice
}

func (s proxy) Swap(i, j int) { s.data.Swap(i, j); s.IntSlice.Swap(i, j) }

func main() {
	data := sort.StringSlice{"zero", "one", "two", "three", "four"}
	Shuffle(data)
	fmt.Println(data)
}
