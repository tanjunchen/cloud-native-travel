package main

/***
"题目：**数组中数字出现的次数II**

[数组中数字出现的次数II](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof)

题目描述：
***/

/**
解法一
说明：哈希值
**/
func singleNumber(nums []int) int {
	res := make(map[int]int, len(nums))
	for _, v := range nums {
		if _, ok := res[v]; ok {
			res[v]++
		} else {
			res[v] = 1
		}
	}
	for k, v := range res {
		if v != 3 {
			return k
		}
	}
	return -1
}

/**
解法二
说明：
**/
func singleNumber2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	count := make([]int, 32)
	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		for j := 0; j < 32; j++ {
			count[j] += tmp & 1
			tmp = tmp >> 1
		}
	}
	res := 0
	for i := 31; i >= 0; i-- {
		res = res << 1
		res = res | count[i]%3
	}
	return res
}
