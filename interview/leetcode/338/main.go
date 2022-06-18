package main

import "fmt"

/***
"题目：**338 比特位计数**

[比特位计数](https://leetcode-cn.com/problems/counting-bits/)

给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

***/
// 方法一：直接计算
func countNum(i int) (count int) {
	for ; i > 0; i = i & (i - 1) {
		count++
	}
	return
}

func countBits(num int) []int {
	bits := make([]int, num+1)
	for i := range bits {
		bits[i] = countNum(i)
	}
	return bits
}

// 低位的动态规划

func countBits2(num int) []int {
	bits := make([]int, num+1)
	for i := 1; i <= num; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}

func main() {
	fmt.Println(countBits(5))
	fmt.Println(countBits2(5))
}
