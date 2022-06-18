package main

/***
"题目：**二进制中1的个数**

[二进制中1的个数](https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num = (num - 1) & num
	}
	return count
}

func hammingWeight2(num uint32) int {
	count := 0
	var flag uint32 = 1
	for flag != 0 {
		if num != 0&flag {
			count++
		}
		flag = flag << 1
	}
	return count
}
