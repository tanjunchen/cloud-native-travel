package main

/***
"题目：**反转链表**

[反转链表](https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof)

题目描述：
***/

/**
解法一
说明：
**/
// 反转链表
// 改变链表的数据结构
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var tmp *ListNode
	for head != nil {
		tmp = &ListNode{Val: head.Val, Next: tmp}
		head = head.Next
	}
	return tmp
}

// 递归操作
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

func reverseList3(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		head, head.Next, newHead = head.Next, newHead, head
	}
	return head
}

// 普遍反转链表的方式
func reverseList4(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
