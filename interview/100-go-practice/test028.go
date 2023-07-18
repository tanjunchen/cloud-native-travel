package main

import (
	"fmt"
)

type People interface {
	Show()
}
type Student struct{}

func (stu *Student) Show() {
}
func live() People {
	var stu *Student
	return stu
}
func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

// 输出值是：BBBBBBB
// *Student 是 nil 的，但是 *Student 实现了 People 接口，接口不为 nil。
