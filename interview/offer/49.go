package main

/***
"题目：**丑数**

[丑数](https://leetcode-cn.com/problems/chou-shu-lcof)

题目描述：
我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

***/

/**
解法一
说明：
**/
func nthUglyNumber(n int) int {
	t1, t2, t3 := 0, 0, 0
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		v1 := dp[t1] * 2
		v2 := dp[t2] * 3
		v3 := dp[t3] * 5
		dp[i] = min(v3, min(v1, v2))
		if dp[i] == v1 {
			t1++
		}
		if dp[i] == v2 {
			t2++
		}
		if dp[i] == v3 {
			t3++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
