package main

/***
"题目：**左旋转字符串**

[左旋转字符串](https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof)

题目描述：
***/

/**
解法一
说明： 如果 n 超过 s 的长度，会有问题
**/
func reverseLeftWords(s string, n int) string {
	return s[n:len(s)] + s[0:n]
}

func reverseLeftWords2(s string, n int) string {
	var str1, str2 string
	for i := 0; i < len(s); i++ {
		if i < n {
			str1 += string(s[i])
		} else {
			str2 += string(s[i])
		}
	}
	return str2 + str1
}

func reverseLeftWords3(s string, n int) string {
	length := len(s)
	n = n % length
	if n == 0 {
		return s
	}
	return string(s[n:length] + s[:n])
}
