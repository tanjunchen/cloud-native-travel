package main

import (
	"fmt"
)

type Student struct {
	name string
}

func main() {
	// map 的 value 本身是不可寻址的，因为 map 中的值会在内存中移动，并且旧的指针地址在 map 改变时会变得无效。
	// cannot assign to struct field m["people"].name in map
	// m := map[string]Student{"people": {"zhoujielun"}}
	// m["people"].name = "wuyanzu"

	m := map[string]*Student{"people": {"zhoujielun"}}
	m["people"].name = "wuyanzu"
	fmt.Println(m["people"]) // "wuyanzu"
}
