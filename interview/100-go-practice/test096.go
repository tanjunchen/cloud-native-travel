package main

import (
	"fmt"
)

type N int

func (n N) value() {
	n++
	fmt.Printf("v:%p,%v\n", &n, n)
}

func (n *N) pointer() {
	*n++
	fmt.Printf("v:%p,%v\n", n, *n)
}

func main() {
	var a N = 25
	// p := &a
	// p1 := &p
	// p1.value()
	// p1.pointer()
	p := a
	p1 := &p
	p1.value()
	p1.pointer()
}

// 不能使用多级指针调用方法。
// ./test096.go:23:5: p1.value undefined (type **N has no field or method value)
// ./test096.go:24:5: p1.pointer undefined (type **N has no field or method pointer)
