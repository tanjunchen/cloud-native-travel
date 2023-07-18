package main

/***
"题目：**从尾到头打印链表**

[从尾到头打印链表](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof)

题目描述：
***/

/**
解法一
说明：
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	start, end := 0, len(values)-1
	for start < end {
		values[start], values[end] = values[end], values[start]
		start++
		end--
	}
	return values
}

/**
解法二
说明：
**/
func reversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var newHead *ListNode
	var res []int
	for head != nil {
		node := head.Next
		head.Next = newHead
		newHead = head
		head = node
	}
	for newHead != nil {
		res = append(res, newHead.Val)
		newHead = newHead.Next
	}
	return res
}

/**
解法三
说明：
**/

func main() {

}
