package main

import (
	"encoding/json"
	"fmt"
)

// 按照 golang 的语法，小写开头的方法、属性是 struct 是私有的
type People struct {
	// 无法正常得到 People 的 name 值
	name string `json:"name"`
}

func main() {
	js := `{
	  "name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}
