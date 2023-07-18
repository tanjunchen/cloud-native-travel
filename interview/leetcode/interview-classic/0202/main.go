package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	newHead := head
	for k > 0 {
		newHead = newHead.Next
		k--
	}
	for newHead != nil {
		head = head.Next
		newHead = newHead.Next
	}
	return head.Val
}

func kthToLast2(head *ListNode, k int) int {
	fast := head
	slow := head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Val
}
