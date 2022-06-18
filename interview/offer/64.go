package main

/***
"题目：**求1+2+…+n**

[求1+2+…+n](https://leetcode-cn.com/problems/qiu-12n-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func sumNums(n int) int {
	return n * (n + 1) / 2
}

func sumNums2(n int) int {
	if n == 1 {
		return 1
	}
	return n + sumNums2(n-1)
}

func sumNums3(n int) int {
	res := 0
	var sum func(int) bool
	sum = func(n int) bool {
		res += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return res
}
