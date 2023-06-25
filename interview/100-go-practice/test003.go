package main

import "fmt"

/**
请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变
量)。
给定一个string，请返回一个string，为翻转后的字符串。保证字符串的⻓度小于等于5000。
*/

func main() {
	s := "abcdefghijklmnopqrstuvwxyz"
	res, ok := reverseString(s)
	fmt.Println(res)
	fmt.Println(ok)
}

func reverseString(str string) (string, bool) {
	s := []rune(str)
	length := len(s)
	if length > 5000 {
		return str, false
	}
	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
	return string(s), true
}
