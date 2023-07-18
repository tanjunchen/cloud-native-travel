package main

/***
"题目：**剪绳子II**

[剪绳子II](https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	if m := 2 - n%3; m == 2 {
		return pow(n / 3)
	} else {
		return pow(n/3-m) * (m + 1) * 2 % 1000000007
	}
}
func pow(n int) int {
	sum := 1
	for i := 0; i < n; i++ {
		sum = sum * 3 % 1000000007
	}
	return sum
}

func cuttingRope3(n int) int {
	if n < 4 {
		return n - 1
	}
	if n == 4 {
		return 4
	}
	res := 1
	for n > 4 {
		res = res * 3 % 1000000007
		n -= 3
	}
	return res * n % 1000000007
}
