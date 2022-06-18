package main

/***
"题目：**缺失的数字**

[缺失的数字](https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof)

题目描述：
***/

/**
解法一
说明：
**/
// 二分查找
func missingNumber(nums []int) int {
	i, j := 0, len(nums)
	for i < j {
		mid := (i + j) >> 1
		if nums[mid] == mid {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

// 标记数组-记录数字
func missingNumber2(nums []int) int {
	flag := make([]bool, len(nums)+1)
	for _, v := range nums {
		flag[v] = true
	}
	for i := 0; i < len(flag); i++ {
		if flag[i] == false {
			return i
		}
	}
	return 0
}

// 数字异或运算
func missingNumber3(nums []int) int {
	var count = len(nums)
	for i := 0; i < len(nums); i++ {
		count ^= i ^ nums[i]
	}
	return count
}

// 总和的差值
func missingNumber4(nums []int) int {
	n := len(nums)
	sum := (n + 1) * n >> 1
	for _, v := range nums {
		sum -= v
	}
	return sum
}
