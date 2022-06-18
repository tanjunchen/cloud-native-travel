package main

import "fmt"

/***
"题目：**1～n整数中1出现的次数**

[1～n整数中1出现的次数](https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof)

题目描述：输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。
***/

/**
解法一
说明：
**/
// 2304
// high cur low
// 23 0 4
// cur = 0   229 + 1     230   high * digit
// cur = 1 234 + 1       235   high * digit + low + 1
// cur 2-9  239 - 0 + 1  240   (high + 1) * digit
func countDigitOne2(n int) int {
	cur := n % 10
	high := n / 10
	low := 0
	digit := 1
	res := 0
	for cur != 0 || high != 0 {
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else {
			res += (high + 1) * digit
		}
		low += cur * digit
		cur = high % 10
		high = high / 10
		digit *= 10
	}
	return res
}

// 超出时间限制
func countDigitOne(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		num := i
		for num > 0 {
			if num%10 == 1 {
				count++
			}
			num = num / 10
		}
	}
	return count
}

/**
解法二
说明：
**/

func main() {
	fmt.Println(countDigitOne(12))
	fmt.Println(countDigitOne(13))
}
