package main

/***
"题目：**打印从1到最大的n位数**

[打印从1到最大的n位数](https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func printNumbers(n int) []int {
	if n < 0 {
		return []int{}
	}
	maxValue := 1
	for i := 0; i < n; i++ {
		maxValue *= 10
	}
	res := make([]int, maxValue-1)
	for i := 0; i < maxValue-1; i++ {
		res[i] = i + 1
	}
	return res
}
