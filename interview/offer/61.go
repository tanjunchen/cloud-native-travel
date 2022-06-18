package main

import (
	"sort"
)

/***
"题目：**扑克牌中的顺子**

[扑克牌中的顺子](https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func isStraight2(nums []int) bool {
	if len(nums) < 5 {
		return false
	}
	sort.Ints(nums)
	value := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
		value += nums[i+1] - nums[i]
	}
	return value < 5
}

/**
解法二
说明：
**/
func isStraight(nums []int) bool {
	if len(nums) < 5 {
		return false
	}
	min, max := 14, 0
	count := 0
	res := make(map[int]bool, 14)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
			continue
		}
		if res[nums[i]] {
			return false
		}
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
		res[nums[i]] = true
	}
	if count == 0 {
		return max-min == 4
	}
	return (max - min) < 5
}

func main() {

}
