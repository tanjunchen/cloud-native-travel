package main

import "fmt"

/***
"题目：**304 二维区域和检索 - 矩阵不可变**

[二维区域和检索 - 矩阵不可变](https://leetcode-cn.com/problems/range-sum-query-2d-immutable/)

给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。

上图子矩阵左上角 (row1, col1) = (2, 1) ，右下角(row2, col2) = (4, 3)，该子矩形内元素的总和为 8。
*/

// 二维矩阵的前缀和
// f(i,j)=f(i−1,j)+f(i,j−1)−f(i−1,j−1)+matrix[i][j]
//
type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{}
	}
	n := len(matrix[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i, row := range matrix {
		sum[i+1] = make([]int, n+1)
		for j, v := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + v
		}
	}
	return NumMatrix{sum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// 套用公式
	return this.sums[row2+1][col2+1] - this.sums[row1][col2+1] - this.sums[row2+1][col1] + this.sums[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {
	num := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	fmt.Println(num.SumRegion(2, 1, 4, 3))
}
