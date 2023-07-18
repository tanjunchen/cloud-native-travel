package main

/***
"题目：**从上到下打印二叉树II**

[从上到下打印二叉树II](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof)

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
func levelOrder2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	var index = 0
	queue = append(queue, root)
	for len(queue) > 0 {
		var temp []*TreeNode
		res = append(res, []int{})
		for i := 0; i < len(queue); i++ {
			res[index] = append(res[index], queue[i].Val)
			if queue[i].Left != nil {
				temp = append(temp, queue[i].Left)
			}
			if queue[i].Right != nil {
				temp = append(temp, queue[i].Right)
			}
		}
		index++
		queue = temp
	}
	return res
}
