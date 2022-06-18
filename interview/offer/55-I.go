package main

/***
"题目：**二叉树的深度**

[二叉树的深度](https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof)

题目描述：
***/

/**
解法一
说明：递归左子树与右子树
**/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
解法二
说明：队列
**/
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root)
	res := 0
	for len(queue) > 0 {
		num := len(queue)
		for i := 0; i < num; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[num:]
		res++
	}
	return res
}
