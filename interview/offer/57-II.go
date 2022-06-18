package main

/***
"题目：**和为s的连续正数序列**

[和为s的连续正数序列](https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof)

题目描述：输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
***/

/**
解法一
说明：滑动窗口的思想
**/

func findContinuousSequence(target int) [][]int {
	if target <= 0 {
		return [][]int{}
	}
	l, r := 1, 1
	sum := 0
	var res [][]int
	for l <= target/2+1 {
		if sum < target {
			sum += r
			r++
		} else if target < sum {
			sum -= l
			l++
		}
		if target == sum {
			var temp []int
			for k := l; k < r; k++ {
				temp = append(temp, k)
			}
			res = append(res, temp)
			sum -= l
			l++
		}
	}
	return res
}
