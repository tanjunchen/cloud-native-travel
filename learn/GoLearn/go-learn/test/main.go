package main

import (
	"fmt"
	"math/rand"
)

func main() {
	value := rand.Int31()
	uniqueName := fmt.Sprintf("%s-%08x", "BB", value)
	fmt.Println(value, uniqueName)
}

// NewString creates a String from a list of values.
func newString(items ...string) String {
	ss := String{}
	ss.Insert(items...)
	return ss
}

// sets.String is a set of strings, implemented via map[string]struct{} for minimal memory consumption.
type String map[string]Empty

// Empty is public since it is used by some internal API objects for conversions between external
// string arrays and internal sets, and conversion logic requires public types today.
type Empty struct{}

// Insert adds items to the set.
func (s String) Insert(items ...string) String {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}
