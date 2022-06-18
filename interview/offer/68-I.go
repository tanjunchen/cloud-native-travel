package main

/***
"题目：**二叉搜索树的最近公共祖先**

[二叉搜索树的最近公共祖先](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof)

题目描述：给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
***/

/**
解法一
说明：
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
class Solution {
public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
TreeNode result = root;
while (true){
if (p.val < result.val && q.val < result.val) {
result = result.left;
} else if (p.val > result.val && q.val > result.val){
result = result.right;
}else{
break;
}
}
return result;
}
}
**/

/**
解法二
说明：
**/

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	result := root
	if result == nil {
		return result
	}
	for {
		if p.Val < result.Val && q.Val < result.Val {
			result = result.Left
		} else if p.Val > result.Val && q.Val > result.Val {
			result = result.Right
		} else {
			break
		}
	}
	return result
}
