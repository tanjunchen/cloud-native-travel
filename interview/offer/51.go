package main

import "fmt"

/***
"题目：**数组中的逆序对**

[数组中的逆序对](https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof)

题目描述：在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
***/

// 暴力法 但是应该会超时
func reversePairs(nums []int) int {
	// 双层循环
	return 0
}

func reversePairs2(nums []int) int {
	count := 0
	calNumWithMerge(nums, &count)
	return count
}

func calNumWithMerge(nums []int, res *int) []int {
	if len(nums) <= 1 {
		return nums
	}
	return merge(calNumWithMerge(nums[:len(nums)/2], res), calNumWithMerge(nums[len(nums)/2:], res), res)
}

func merge(l []int, r []int, res *int) []int {
	lLen, rLen := len(l), len(r)
	var list []int
	var i, j int
	for i < lLen && j < rLen {
		if l[i] <= r[j] {
			list = append(list, l[i])
			i++
		} else {
			list = append(list, r[j])
			j++
			*res = *res + lLen - i
		}
	}
	if i < lLen {
		list = append(list, l[i:]...)
	}
	if j < rLen {
		list = append(list, r[j:]...)
	}
	return list
}

/**
解法一
说明：
**/

func reversePairs3(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	var tmp []int
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			cnt += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += end - (mid + 1) + 1
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
	return cnt
}

func main() {
	fmt.Println(reversePairs2([]int{7, 5, 6, 4}))
}
