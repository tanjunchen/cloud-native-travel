package main

import "fmt"

/***
"题目：**697. 数组的度**

[数组的度](https://leetcode-cn.com/problems/degree-of-an-array/)

题目描述：给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。

你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

***/

/**
解法一
说明：循环 + 哈希
**/

type Info struct {
	len   int
	start int
	count int
}

func findShortestSubArray(nums []int) int {
	data := map[int]*Info{}
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		value, ok := data[v]
		if ok {
			value.count++
			value.len = i - value.start
		} else {
			value = &Info{
				len:   0,
				start: i,
				count: 1,
			}
		}
		data[v] = value
	}

	max := 0
	shortLen := len(nums)
	for _, v := range data {
		if v.count > max {
			max = v.count
			shortLen = v.len
		} else if v.count == max && v.len <= shortLen {
			max = v.count
			shortLen = v.len
		}
	}
	return shortLen + 1
}

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
}
