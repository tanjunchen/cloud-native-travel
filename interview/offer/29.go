package main

import "fmt"

/***
"题目：**顺时针打印矩阵**

[顺时针打印矩阵](https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof)

题目描述：
***/

/**
解法一
说明：左到右 上到下 右到左 下到上
**/
func spiralOrder(matrix [][]int) []int {
	// 1. 条件判断 是否为空
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	// 2. 行列
	row, column := len(matrix), len(matrix[0])
	left, top, bottom, right := 0, 0, row-1, column-1
	index := 0
	order := make([]int, row*column)
	for left <= right && top <= bottom {
		// 3. 左到右
		for c := left; c <= right; c++ {
			order[index] = matrix[top][c]
			index++
		}
		// 4. 上到下
		for r := top + 1; r <= bottom; r++ {
			order[index] = matrix[r][right]
			index++
		}
		if left < right && top < bottom {
			for c := right - 1; c > left; c-- {
				order[index] = matrix[bottom][c]
				index++
			}
			for r := bottom; r > top; r-- {
				order[index] = matrix[r][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}

func main() {
	res := spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println(res)
}
