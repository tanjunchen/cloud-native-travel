package main

import "fmt"

/***
"题目：**303 区域和检索 - 数组不可变**

[区域和检索 - 数组不可变](https://leetcode-cn.com/problems/range-sum-query-immutable/)

给定一个整数数组  nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点。

实现 NumArray 类：

NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int i, int j) 返回数组 nums 从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点（也就是 sum(nums[i], nums[i + 1], ... , nums[j])）

***/
/**
输入：
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
输出：
[null, 1, -1, -3]
*/
/**
暴力法
优化 内层循环累加求和时：上轮迭代求了 i 到 j - 1 的和，这轮就没必要从头求 i 到 j 的和。
前缀和
*/

type NumArray struct {
	preSum []int
	len    int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{
			preSum: make([]int, 0),
			len:    0,
		}
	}
	na := NumArray{
		preSum: make([]int, len(nums)),
		len:    len(nums),
	}
	na.preSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		na.preSum[i] = na.preSum[i-1] + nums[i]
	}
	return na
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		if this.len == 0 {
			return 0
		}
		return this.preSum[j]
	}
	// nums[i]+…+nums[j]=preSum[j]−preSum[i−1]
	// sumRange(i,j)=preSum[j]−preSum[i−1]
	return this.preSum[j] - this.preSum[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

/**
前缀和简化版
*/
type ArrayNum struct {
	preSum []int
}

func NewArrayNum(nums []int) ArrayNum {
	na := ArrayNum{preSum: make([]int, len(nums)+1)}
	na.preSum[0] = 0
	for i := 0; i < len(nums); i++ {
		na.preSum[i+1] = na.preSum[i] + nums[i]
	}
	return na
}
func (this ArrayNum) SumRange(i int, j int) int {
	return this.preSum[j+1] - this.preSum[i]
}

func main() {
	obj := NewArrayNum([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}
