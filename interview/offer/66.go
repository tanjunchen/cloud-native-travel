package main

/***
"题目：**构建乘积数组**

[构建乘积数组](https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func constructArr2(a []int) []int {
	// 所有的乘积除以 index 的值 但是不能使用除法
	// 某个元素的前缀积与后缀积的乘积
	length := len(a)
	if length <= 1 {
		return a
	}
	res := make([]int, length)
	res[0] = 1
	for i := 1; i < length; i++ {
		res[i] = res[i-1] * a[i-1]
	}
	res[0] = a[length-1]
	for i := length - 2; i >= 1; i-- {
		res[i] *= res[0]
		res[0] *= a[i]
	}
	return res
}

/**
解法二
说明：前缀积与后缀积之积
**/
func constructArr(a []int) []int {
	// 所有的乘积除以 index 的值 但是不能使用除法
	// 某个元素的前缀积与后缀积的乘积
	length := len(a)
	if length <= 1 {
		return a
	}
	l, r := make([]int, length), make([]int, length)
	l[0], r[length-1] = 1, 1
	for i := 1; i < length; i++ {
		l[i] = l[i-1] * a[i-1]
	}
	for i := length - 2; i >= 0; i-- {
		r[i] = r[i+1] * a[i+1]
	}

	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = l[i] * r[i]
	}
	return result
}
