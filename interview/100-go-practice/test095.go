package main

import (
	"fmt"
)

type T struct {
	n int
}

func (t *T) Set(n int) {
	t.n = n
}
func getT() T {
	return T{}
}
func main() {
	// getT().Set(1)
	t := getT()
	t.Set(2)
	fmt.Println(t.n)
}

// 1.直接返回的 T{} 不可寻址；
// 2.不可寻址的结构体不能调用带结构体指针接收者的方法
// ./test095.go:14:9: cannot call pointer method Set on T
