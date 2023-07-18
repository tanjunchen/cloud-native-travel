package main

import "fmt"

/***
"题目：**503. 下一个更大元素 II**

[下一个更大元素 II](https://leetcode-cn.com/problems/next-greater-element-ii/)

给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。
数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

***/

func main() {
	// 2 -1 2
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	// 2 3 4 -1
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
}

// 单调栈的解法
func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var stack []int
	for i := 0; i < 2*n-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return ans
}
