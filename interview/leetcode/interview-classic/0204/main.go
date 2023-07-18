package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	first := &ListNode{}
	firstHead := first

	second := &ListNode{}
	secondHead := second

	for head != nil {
		if head.Val < x {
			first.Next = head
			first = first.Next
		} else {
			second.Next = head
			second = second.Next
		}
		head = head.Next
	}

	second.Next = nil
	first.Next = secondHead.Next
	return firstHead.Next
}
