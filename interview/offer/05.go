package main

import "strings"

/***
"题目：**替换空格**

[替换空格](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func replaceSpace(s string) string {
	var str strings.Builder
	for _, i := range s {
		if string(i) == " " {
			str.WriteString("%20")
		} else {
			str.WriteString(string(i))
		}
	}
	return str.String()
}

/**
解法二
说明：
**/

func replaceSpace2(s string) string {
	var str string = ""
	for _, v := range s {
		if v == ' ' {
			str += "%20"
		} else {
			str += string(v)
		}
	}
	return str
}

/**
解法三
说明：
**/

func main() {

}
