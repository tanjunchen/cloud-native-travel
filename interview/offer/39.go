package main

import "sort"

/***
"题目：**数组中出现次数超过一半的数字**

[数组中出现次数超过一半的数字](https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof)

题目描述：
***/

/**
解法一
说明：众数投票法
**/
func majorityElement(nums []int) int {
	// map
	// 排序 中间值
	// 投票方式
	if len(nums) == 0 {
		return -1
	}
	count := 1
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			res = nums[i]
		}
		if res != nums[i] {
			count--
		} else {
			count++
		}
	}
	return res
}

/**
解法二
说明：
**/
func majorityElement2(nums []int) int {
	m := make(map[int]int)
	t := len(nums) / 2
	for _, v := range nums {
		if m[v] >= t {
			return v
		}
		m[v]++
	}
	return -1
}

/**
解法三
说明：
**/
func majorityElement3(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
