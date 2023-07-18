package main

import (
	"fmt"
	"reflect"
)

type strTemp struct {
	newStr string
}
type str struct {
	a  string
	st *strTemp
}

func main() {
	var p *str
	check(p)
}

func check(i interface{}) {
	fmt.Printf("value=%v type=%t\n", i, i)
	if i == nil {
		fmt.Printf("%s is nil ...", i)
	}
	if reflect.ValueOf(i).IsNil() {
		fmt.Println("value i is nil")
	}
	fmt.Println("Hello Golang")
}
