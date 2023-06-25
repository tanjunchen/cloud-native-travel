package main

import (
	"fmt"
)

func main() {
	s := make([]int, 3, 9)
	fmt.Println(len(s))
	s2 := s[4:8]
	fmt.Println(len(s2))
	// 代码没问题，输出 3 4
}
