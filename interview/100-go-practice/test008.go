package main

import (
	"fmt"
)

type student struct {
	Name string
}

// golang中有规定，switch type 的case T1 ，类型列表只有一个，那么v := m.(type) 中的v 的类型就是T1类
型。
// func zhoujielun(v interface{}) {
// 	switch msg := v.(type) {
// 	case *student, student:
// 		msg.Name
// 	}
// }

func zhoujielun2(v interface{}) {
	switch msg := v.(type) {
	case *student:
		name := msg.Name
		fmt.Println("*student===>", name)
	case student:
		name := msg.Name
		fmt.Println("student===>", name)
	}
}

func main() {
	zhoujielun2(student{Name: "liangjian"})
}
