package main

/***
"题目：**圆圈中最后剩下的数字**

[圆圈中最后剩下的数字](https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof)

题目描述：
***/

/**
解法一
说明：迭代法 非递归
**/
func lastRemaining(n int, m int) int {
	res := 0
	for i := 2; i <= n; i++ {
		res = (m + res) % i
	}
	return res
}

/**
解法二
说明：递归
**/
func lastRemaining2(n int, m int) int {
	if n == 1 {
		return 0
	}
	return (m + lastRemaining(n-1, m)) % n
}
