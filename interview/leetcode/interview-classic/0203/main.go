package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	*node = *node.Next
}

func deleteNode2(node *ListNode) {
	//*node = *node.Next
	for node != nil {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
		}
		node = node.Next
	}
}
