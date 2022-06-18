package main

/***
"题目：**矩阵中的路径**

[矩阵中的路径](https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	if len(word) == 0 {
		return true
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				w := 0
				if dfs(board, i, j, w, word) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, k int, word string) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	temp := board[i][j]
	board[i][j] = '#'
	res := dfs(board, i, j+1, k+1, word) || dfs(board, i, j-1, k+1, word) || dfs(board, i-1, j, k+1, word) || dfs(board, i+1, j+1, k+1, word)
	board[i][j] = temp
	return res
}
