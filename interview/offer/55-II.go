package main

import "math"

/***
"题目：**平衡二叉树**

[平衡二叉树](https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof)

题目描述：
***/

/**
解法一
说明：
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return dfsTree(root) != -1
}

func dfsTree(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	l := dfsTree(root.Left)
	if l == -1 {
		return -1
	}
	r := dfsTree(root.Right)
	if r == -1 {
		return -1
	}
	if math.Abs(l-r) > 1 {
		return -1
	}
	return math.Max(l, r) + 1
}
