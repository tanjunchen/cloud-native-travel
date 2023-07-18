package main

import (
	"fmt"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	// golang 语言中没有继承概念，只有组合，也没有虚方法，更没有重载。
	// 因此，*Teacher 的 ShowB 不会覆写被组合的 People 的方法。
	t.ShowA()
}
