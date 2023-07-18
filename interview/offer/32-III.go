package main

/***
"题目：**从上到下打印二叉树III**

[从上到下打印二叉树III](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof)

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
func levelOrder4(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	var direction = true
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		list := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i]
			if direction {
				list[i] = node.Val
			} else {
				list[size-i-1] = node.Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
		direction = !direction
		res = append(res, list)
	}
	return res
}
