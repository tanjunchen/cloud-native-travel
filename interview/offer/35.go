package main

/***
"题目：**复杂链表的复制**

[复杂链表的复制](https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof)

题目描述：请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。
***/

/**
解法一
说明：
**/
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	p := head.Next
	pre := &newHead
	connection := make(map[*Node]*Node)
	connection[head] = pre
	for p != nil {
		newNode := &Node{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}
		pre.Next = newNode
		connection[p] = newNode
		p = p.Next
		pre = pre.Next
	}
	p = head
	newP := &newHead
	for p != nil {
		if p.Random != nil {
			newP.Random = connection[p.Random]
		}
		p = p.Next
		newP = newP.Next
	}
	return &newHead
}

/**
解法二
说明：
**/
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	p := head
	for p != nil {
		newNode := &Node{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}
		newNode.Next = p.Next
		p.Next = newNode
		p = p.Next.Next
	}
	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}
	newHead := head.Next
	oldHead := head
	p = newHead
	for p.Next != nil {
		oldHead.Next = oldHead.Next.Next
		p.Next = p.Next.Next
		oldHead = oldHead.Next
		p = p.Next
	}
	oldHead.Next = nil
	return newHead
}

func main() {

}
