package main

import (
	"fmt"
)

func main() {
	// [0 0 0 0 0 1 2 3]
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
