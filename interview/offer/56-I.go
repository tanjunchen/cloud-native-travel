package main

/***
"题目：**数组中数字出现的次数**

[数组中数字出现的次数](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof)

题目描述：一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
***/

/**
解法一
说明：因为有要求：时间复杂度与空间复杂度
**/

/**
解法二
说明：map 哈希值的方式不太合适
**/
func singleNumbers(nums []int) (res []int) {
	maps := make(map[int]int)
	for _, value := range nums {
		maps[value]++
	}
	for key, value := range maps {
		if value == 1 {
			res = append(res, key)
		}
	}
	return res
}

/**
解法二
说明：
**/
func singleNumbers2(nums []int) []int {
	diff := 0
	for i := 0; i < len(nums); i++ {
		diff ^= nums[i]
	}
	diff = -diff & diff
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		if diff&nums[i] == 0 {
			res[0] ^= nums[i]
		} else {
			res[1] ^= nums[i]
		}
	}
	return res
}
