package main

/***
"题目：**数值的整数次方**

[数值的整数次方](https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func myPow(x float64, n int) float64 {
	sum := 1.0
	flag := false
	if n < 0 {
		n = -n
		flag = true
	}
	for i := 0; i < n; i++ {
		sum *= x
	}
	if flag {
		sum = 1.0 / sum
	}
	return sum
}

// 迭代法
func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	res := 1.0
	for n >= 1 {
		if n&1 == 1 {
			res *= x
			n--
		} else {
			x *= x
			n = n >> 1
		}
	}
	return res
}

// 递归法
func myPow3(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	temp := myPow(x, n/2)
	if n%2 == 0 {
		return temp * temp
	}
	return x * temp * temp
}
