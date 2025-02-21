package main

type TreeNode struct {
	Val  int
	Next *TreeNode
}

func merge(list1, list2 *TreeNode) *TreeNode {
	dummy := *&TreeNode{
		Val: -1,
	}
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			dummy.Next = list1
			list1 = list1.Next
		} else {
			dummy.Next = list2
			list2 = list2.Next
		}
	}
	for list1 != nil {
		dummy.Next = list1
		list1 = list1.Next
	}

	for list2 != nil {
		dummy.Next = list2
		list2 = list2.Next
	}
	return dummy.Next
}

func main() {

}
