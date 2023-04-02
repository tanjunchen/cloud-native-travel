package main

import (
	"fmt"
	"strings"
)

func test01() {

	a := "gopher"
	b := "hello world"

	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Compare(a, a))
	fmt.Println(strings.Compare(b, a))

	fmt.Println(strings.EqualFold("GO", "go"))
	fmt.Println(strings.EqualFold("壹", "一"))

	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))
}

func main() {
	test01()
}
