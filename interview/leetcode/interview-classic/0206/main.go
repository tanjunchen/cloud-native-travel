package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	l := 0
	r := len(res) - 1
	for l < r {
		if res[l] != res[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
}
