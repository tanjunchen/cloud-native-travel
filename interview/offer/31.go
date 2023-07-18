package main

/***
"题目：**栈的压入、弹出序列**

[栈的压入、弹出序列](https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for popped[0] == stack[len(stack)-1] {
			popped = popped[1:]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
		}
	}
	return len(stack) == 0
}
