package main

import "fmt"
import "strconv"

// type ListNode struct {
//     Val int
//     Next *ListNode
// }

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(letterCombinations("23"))
}

// 2 abc
// 3 def
// 4 ghi
// 5 jkl
// 6 mno
// 7 pqrs
// 8 tuv
// 9 wxyz

func minDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	queue := []*TreeNode{root}
	count := []int{1}
	for i:=0;i< len(queue);i++{
		node := queue[i]
		depth := count[i]
		if node.Left == nil && node.Right == nil{
			return depth
		}
		if node.Left != nil{
			queue = append(queue, node.Left)
			count = append(count, depth+1)
		}
		if node.Right != nil{
			queue = append(queue, node.Right)
			count = append(count, depth+1)
		}
	}
	return 0
}

func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
	if root == nil{
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		tmp := make([]int,n)
		for i:=0;i<n;i++{
			node := queue[0]
			queue = queue[1:]
			tmp[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res,tmp)
	}
	return res
}

func max(x,y int) int{
	if x > y{
		return x
	}
	return y
}

func min(x,y int) int{
	if x < y{
		return x
	}
	return y
}
