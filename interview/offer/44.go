package main

import "strconv"

/***
"题目：**数字序列中某一位的数字**

[数字序列中某一位的数字](https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof)

题目描述：数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

请写一个函数，求任意第n位对应的数字。
***/

/**
解法一
说明：
**/
func findNthDigit(n int) int {
	start := 1
	count := 9
	digst := 1
	for n-count > 0 {
		n -= count
		start *= 10
		digst += 1
		count = start * 9 * digst
	}
	num := start + (n-1)/digst
	res := strconv.Itoa(num)[(n-1)%digst] - '0'
	return int(res)
}

func main() {

}
