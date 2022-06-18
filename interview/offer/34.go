package main

/***
"题目：**二叉树中和为某一值的路径**

[二叉树中和为某一值的路径](https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof)

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
func pathSum(root *TreeNode, sum int) [][]int {
	var data [][]int
	if root == nil {
		return data
	}
	dfs34(root, sum, []int{}, &data)
	return data
}

func dfs34(root *TreeNode, sum int, arr []int, res *[][]int) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)
	if root.Val == sum && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*res = append(*res, tmp)
	}
	dfs34(root.Left, sum-root.Val, arr, res)
	dfs34(root.Right, sum-root.Val, arr, res)
	arr = arr[:len(arr)-1]
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
func pathSum2(root *TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var path []int
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		target -= root.Val
		if 0 == target && root.Left == nil && root.Right == nil {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		dfs(root.Left, target)
		dfs(root.Right, target)
		path = path[0 : len(path)-1]
	}
	dfs(root, sum)
	return res
}

func main() {

}
