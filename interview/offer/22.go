package main

/***
"题目：**链表中倒数第k个节点**

[链表中倒数第k个节点](https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof)

题目描述：
***/

/**
解法一
说明：
**/

// 链表中倒数第 k 个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	count := 0
	for head != nil {
		head = head.Next
		count++
	}
	for i := 0; i < count-k; i++ {
		cur = cur.Next
	}
	return cur
}
