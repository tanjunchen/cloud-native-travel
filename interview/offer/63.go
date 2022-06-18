package main

/***
"题目：**股票的最大利润**

[股票的最大利润](https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var min, max int
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			min = prices[0]
			continue
		}
		if prices[i]-min > max {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return max
}

func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min, max := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if (prices[i] - min) > max {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return max
}
