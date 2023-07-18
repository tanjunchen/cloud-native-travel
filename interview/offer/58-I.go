package main

import "strings"

/***
"题目：**翻转单词顺序**

[翻转单词顺序](https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func reverseWords(s string) string {
	strings.TrimSpace(s)
	strs := strings.Split(s, " ")
	var res string
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] == "" {
			continue
		}
		res += strs[i] + " "
	}
	return strings.TrimSpace(res)
}

func reverseWords2(s string) string {
	var strs []string
	var fast int
	for fast < len(s) {
		var str string
		//找到第一个不是空格的
		for fast < len(s) && s[fast] == ' ' {
			fast++
		}
		//找到下一个空格
		for fast < len(s) && s[fast] != ' ' {
			str += string(s[fast])
			fast++
		}
		if len(str) > 0 {
			strs = append(strs, str)
		}
	}
	for i := 0; i < len(strs)/2; i++ {
		strs[i], strs[len(strs)-i-1] = strs[len(strs)-i-1], strs[i]
	}
	return strings.Join(strs, " ")
}

func reverseWords3(s string) string {
	strList := strings.Split(s, " ")
	var res []string
	for i := len(strList) - 1; i >= 0; i-- {
		str := strings.TrimSpace(strList[i])
		if len(str) > 0 {
			res = append(res, strList[i])
		}
	}
	return strings.Join(res, " ")
}
