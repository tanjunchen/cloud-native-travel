package main

/***
"题目：**二叉搜索树的后序遍历序列**

[二叉搜索树的后序遍历序列](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func verifyPostorder(postorder []int) bool {
	if len(postorder) < 2 {
		return true
	}
	size := len(postorder) - 1
	for size != 0 {
		cur := 0
		for postorder[cur] < postorder[size] {
			cur++
		}
		for postorder[cur] > postorder[size] {
			cur++
		}
		if cur != size {
			return false
		}
		size--
	}
	return true
}

/**
解法二
说明：
**/
func verifyPostorder2(postorder []int) bool {
	if len(postorder) < 2 {
		return true
	}
	return judge(postorder, 0, len(postorder)-1)
}
func judge(postorder []int, start, end int) bool {
	if start >= end {
		return true
	}
	i := 0
	for i = start; i < end; i++ {
		if postorder[i] > postorder[end] {
			break
		}
	}
	for j := i; j < end; j++ {
		if postorder[j] < postorder[end] {
			return false
		}
	}
	return judge(postorder, start, i-1) && judge(postorder, i, end-1)
}
