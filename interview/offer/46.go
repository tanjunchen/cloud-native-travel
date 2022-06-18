package main

import "fmt"

/***
"题目：**把数字翻译成字符串**

[把数字翻译成字符串](https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof)

题目描述：给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。
一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

***/

/**
解法一
说明：
**/

func translateNum2(num int) int {
	// 前一个数字
	pn := -1
	// 当前数字
	cn := 0
	// 默认值
	ans := 1
	// 前一位的数量(i-2)
	last := ans
	for num > 0 {
		cn = num % 10
		num = num / 10
		if cn != 0 && pn >= 0 && cn*10+pn <= 25 {
			ans, last = ans+last, ans
		} else {
			last = ans
		}
		pn = cn
	}
	return ans
}

/**
解法二
说明：
**/
func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	var res int

	if num%100 <= 25 && num%100 > 9 {
		res += translateNum(num / 100)
		res += translateNum(num / 10)
	} else {
		res += translateNum(num / 10)
	}

	return res
}

func main() {
	fmt.Println(translateNum(12258))
	fmt.Println(translateNum2(12258))
}
