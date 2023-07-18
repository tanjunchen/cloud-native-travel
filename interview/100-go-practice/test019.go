package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

// golang 的 for ... range 语法中，stu 变量会被复用，每次循环会将集合中的值复制给这个变量
func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
		fmt.Println("Name:", stu.Name, " Age:", stu.Age)
	}
	for k, v := range m {
		fmt.Println("k:", k, " v:", v)
	}
}

func main() {
	pase_student()
}
