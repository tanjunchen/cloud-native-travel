package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	// s := new(Show)
	// s.Param["x"] = 4
	s := Show{Param: make(map[string]interface{}, 0)}
	s.Param["x"] = 4
	fmt.Println(s.Param)
}
