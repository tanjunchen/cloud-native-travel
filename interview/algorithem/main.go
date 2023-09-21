package main

import "fmt"
import "strconv"

// type ListNode struct {
//     Val int
//     Next *ListNode
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	flag := map[int]bool{}

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			parent[root.Left.Val] = root
			dfs(root.Left)
		}
		if root.Right != nil {
			parent[root.Right.Val] = root
			dfs(root.Right)
		}
	}
	dfs(root)
	for p != nil {
		flag[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if flag[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

func reverse(tmp []string) []string {
	for i := 0; i < len(tmp)/2; i++ {
		tmp[i], tmp[len(tmp)-i-1] = tmp[len(tmp)-i-1], tmp[i]
	}
	return tmp
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
