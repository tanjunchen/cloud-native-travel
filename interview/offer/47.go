package main

/***
"题目：**礼物的最大价值**

[礼物的最大价值](https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof)

题目描述：
***/

/**
解法一
说明：dp 动态规划的问题
**/
func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += max2(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
