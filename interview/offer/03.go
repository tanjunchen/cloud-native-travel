package main

import "sort"

/***
"题目：**数组中重复的数字**

[数组中重复的数字](https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof)

题目描述：
***/

/**
解法一
说明：哈希
**/
func findRepeatNumber(nums []int) int {
	if nums == nil {
		return -1
	}
	repeatedMap := make(map[int]int)
	for _, i := range nums {
		if _, ok := repeatedMap[i]; ok {
			return i
		}
		repeatedMap[i] = 1
	}
	return -1
}

/**
解法二
说明：
**/
func findRepeatNumber2(nums []int) int {
	if nums == nil {
		return -1
	}
	length := len(nums)
	for i := 0; i < length; i++ {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	return -1
}

/**
解法三
说明：先排序, 后查找
**/
func findRepeatNumber3(nums []int) int {
	if nums == nil {
		return -1
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

func main() {

}
