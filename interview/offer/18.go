package main

/***
"题目：**删除链表的节点**

[删除链表的节点](https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof)

题目描述：
***/

/**
解法一
说明：
**/

// 单指针方式
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}
	pre := head
	for pre.Next != nil && pre.Next.Val != val {
		pre = pre.Next
	}
	if pre.Next != nil {
		pre.Next = pre.Next.Next
	}
	return head
}

// 递归方式
func deleteNode2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	head.Next = deleteNode2(head.Next, val)
	return head
}

// 双指针方式
func deleteNode3(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	pre, cur := head, head
	for cur != nil && cur.Val != val {
		pre, cur = cur, cur.Next
	}
	if cur != nil {
		pre.Next = cur.Next
	}
	return head
}
