package main

/***
"题目：**旋转数组的最小数字**

[旋转数组的最小数字](https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func minArray(numbers []int) int {
	length := len(numbers)
	if length == 0 {
		return -1
	}
	min := numbers[0]
	for i := 1; i < length; i++ {
		if min > numbers[i] {
			min = numbers[i]
		}
	}
	return min
}

/**
解法二
说明：
**/
func minArray2(numbers []int) int {
	length := len(numbers)
	if length == 0 {
		return -1
	}
	l, r := 0, length-1
	for l < r {
		mid := l + (r-l)>>2
		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else {
			r--
		}
	}
	return numbers[l]
}
