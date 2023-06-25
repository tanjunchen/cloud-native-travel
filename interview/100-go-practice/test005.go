package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**
请编写一个方法，将字符串中的空格全部替换为“%20”。 假定该字符串有足够的空间存放新增的字符，并且知道字符串的真实⻓度(小于等于1000)，同时保证字符串由【大 小写的英文字母组成】。
给定一个string为原始的串，返回替换后的string。
*/

func main() {
	// 使用golang内置方法 unicode.IsLetter 判断字符是否是字母，之后使用 strings.Replace 来替换空格
	fmt.Println(replaceSpaces("hello world"))
}

func replaceSpaces(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false
	}
	for _, c := range s {
		if string(c) != " " && !unicode.IsLetter(c) {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}
