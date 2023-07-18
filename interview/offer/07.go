package main

/***
"题目：**重建二叉树**

[重建二叉树](https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof)

题目描述：
***/

/**
解法一
说明：
**/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var left, right []int
	for k, _ := range inorder {
		if preorder[0] == inorder[k] {
			left = inorder[0:k]
			right = inorder[k+1 : len(inorder)]
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:len(left)+1], left),
		Right: buildTree(preorder[len(left)+1:], right),
	}
}

func main() {

}
