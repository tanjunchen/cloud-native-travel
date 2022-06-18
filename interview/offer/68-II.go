package main

/***
"题目：**二叉树的最近公共祖先**

[二叉树的最近公共祖先](https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof)

题目描述：给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。
***/

/**
解法二
说明：java 语言的解法
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
if (root == null || p == root || q == root){
return root;
}
TreeNode l = lowestCommonAncestor(root.left, p , q);
TreeNode r = lowestCommonAncestor(root.right, p, q);
return l == null ? r : (r == null ? l : root);
}
}
**/

/**
解法三
说明：go 语言的解法
class Solution {
public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
if (root == null || p == root || q == root){
return root;
}
TreeNode l = lowestCommonAncestor(root.left, p , q);
TreeNode r = lowestCommonAncestor(root.right, p, q);
return l == null ? r : (r == null ? l : root);
}
}
**/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Left, p, q)
	if l != nil && r != nil {
		return root
	}
	if l == nil {
		return r
	}
	return l
}
