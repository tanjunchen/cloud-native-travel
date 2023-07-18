package main

import (
	"fmt"
	"strings"
)

/**
请实现一个算法，确定一个字符串的所有字符【是否全都不同】。
这里我们要求【不允许使用额外的存储结构】。
给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。
保证字符串中的字符为【ASCII字符】。字符串的⻓度小于等于【3000】。
*/

func main() {
	fmt.Println(isUniqueString("aaabbb"))
	fmt.Println(isUniqueString("ab"))
	fmt.Println(isUniqueString2("aaabbb"))
	fmt.Println(isUniqueString2("ab"))
}

func isUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, c := range s {
		if c > 127 {
			return false
		}
		if strings.Count(s, string(c)) > 1 {
			return false
		}
	}
	return true
}

func isUniqueString2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for k, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}
