package main

/***
"题目：**二叉搜索树的第 k 大节点**

[二叉搜索树的第k大节点](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof)

题目描述：
***/

/**
解法一
说明：利用中序遍历的顺序性
**/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
	// 二叉搜索树
	if root == nil || k < 0 {
		return -1
	}
	var res []int
	dfsNode(root, &res)
	return res[len(res)-k]
}

func dfsNode(node *TreeNode, res *[]int) {
	if node != nil {
		dfsNode(node.Left, res)
		*res = append(*res, node.Val)
		dfsNode(node.Right, res)
	}
}

/**
解法二
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
var skip int
var res int

func kthLargest2(root *TreeNode, k int) int {
	skip = k
	res = 0
	dfsNode2(root)
	return res
}

func dfsNode2(root *TreeNode) {
	if root != nil {
		dfsNode2(root.Right)
		skip--
		if skip == 0 {
			res = root.Val
			return
		}
		dfsNode2(root.Left)
	}
}

/**
解法三
说明：
**/
