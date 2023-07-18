package main

/***
"题目：**不用加减乘除做加法**

[不用加减乘除做加法](https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof)

题目描述：
***/

/**
解法一
说明：逻辑异或操作
**/
func add(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}
