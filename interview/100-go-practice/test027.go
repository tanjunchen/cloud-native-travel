package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	// var peo People = &Student{} ok
	var peo People = Student{} // error Student 与 *Student 是两种类型
	think := "bitch"
	fmt.Println(peo.Speak(think))
	// ./test027.go:24:19: cannot use Student{} (value of type Student) as type People in variable declaration:
	// Student does not implement People (Speak method has pointer receiver)
}
