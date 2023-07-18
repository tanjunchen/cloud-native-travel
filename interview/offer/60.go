package main

/***
"题目：**n个骰子的点数**

[n个骰子的点数](https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof)

题目描述：把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。

 

你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。
***/

/**
解法一
说明：
**/
import "math"

func dicesProbability(n int) []float64 {
	sum := math.Pow(6.0, float64(n))
	dp := make([]int, 6*n+1)
	var res []float64
	for i := 1; i <= 6; i++ {
		dp[i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 6 * i; j >= i; j-- {
			dp[j] = 0
			for k := 1; k <= 6; k++ {
				if j-k >= i-1 {
					dp[j] += dp[j-k]
				}
			}
		}
	}
	for k := n; k <= 6*n; k++ {
		res = append(res, float64(dp[k])/sum)
	}
	return res
}
