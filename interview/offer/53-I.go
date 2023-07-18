package main

/***
"题目：**在排序数组中查找数字**

[在排序数组中查找数字](https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof)

题目描述：
***/

/**
解法一
说明：

**/
func search(nums []int, target int) int {
	// 直接一次遍历查找
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}

func search2(nums []int, target int) int {
	// 直接一次遍历查找
	res := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := res[nums[i]]; ok {
			res[nums[i]]++
		} else {
			res[nums[i]] = 1
		}
	}
	return res[target]
}

func search3(nums []int, target int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}

	left, right, mid, times := 0, n-1, 0, 0

	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			times++
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 没有找到目标值
	if left > right {
		return 0
	}

	// 向左查找
	for i := mid - 1; i >= 0; i-- {
		if nums[i] == target {
			times++
		} else {
			break
		}
	}

	// 向右查找
	for i := mid + 1; i < n; i++ {
		if nums[i] == target {
			times++
		} else {
			break
		}
	}

	return times
}
