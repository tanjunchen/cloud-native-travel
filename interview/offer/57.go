package main

/***
"题目：**和为s的两个数字**

[和为s的两个数字](https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] == target {
			return []int{nums[i], nums[j]}
		}
		if nums[j]+nums[i] > target {
			j--
		}
		if nums[i]+nums[j] < target {
			i++
		}
	}
	return []int{}
}
