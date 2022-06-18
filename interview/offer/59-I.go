package main

/***
"题目：**滑动窗口的最大值**

[滑动窗口的最大值](https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		return nil
	}
	var maxNums []int
	max := -1
	for i := 0; i <= len(nums)-k; i++ {
		l := i
		r := i + k - 1
		if max == -1 || max == l-1 {
			max = getMax(nums, l, r)
		} else {
			if nums[r] > nums[max] {
				max = r
			}
		}
		maxNums = append(maxNums, nums[max])
	}
	return maxNums
}

func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		return nil
	}
	var maxNums []int
	var max int
	for i, j := 0, k-1; j >= 0 && j < len(nums); j++ {
		if i == 0 || max == nums[i-1] {
			max = nums[i]
			for t := j; t > i; t-- {
				if max < nums[t] {
					max = nums[t]
				}
			}
		} else {
			if nums[j] > max {
				max = nums[j]
			}
		}
		maxNums = append(maxNums, max)
		i++
	}
	return maxNums
}

func getMax(nums []int, l, r int) int {
	for l < r {
		if nums[l] > nums[r] {
			r--
		} else {
			l++
		}
	}
	return l
}
