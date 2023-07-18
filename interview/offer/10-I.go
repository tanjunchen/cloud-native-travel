package main

/***
"题目：**斐波那契数列**

[斐波那契数列](https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1)%1000000007 + fib(n-2)%1000000007
}

func fib2(n int) int {
	f1, f2 := 0, 1
	for i := 0; i < n; i++ {
		f1, f2 = f2, (f1+f2)%1000000007
	}
	return f1
}

func fib3(n int) int {
	if n < 1 {
		return n
	}
	values := make([]int, n+1)
	values[0] = 0
	values[1] = 1
	for i := 2; i <= n; i++ {
		values[i] = values[i-1]%1000000007 + values[i-2]%1000000007
	}
	return values[n] % 1000000007
}