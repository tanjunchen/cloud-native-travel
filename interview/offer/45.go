package main

import (
	"bytes"
	"fmt"
	"sort"
)

/***
"题目：**把数组排成最小的数**

[把数组排成最小的数](https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compare(nums[i], nums[j])
	})
	var res bytes.Buffer
	for i := 0; i < len(nums); i++ {
		res.WriteString(fmt.Sprintf("%d", nums[i]))
	}
	return res.String()
}

func compare(a int, b int) bool {
	if fmt.Sprintf("%d%d", a, b) < fmt.Sprintf("%d%d", b, a) {
		return true
	}
	return false
}

/**
解法二
说明：
**/

/**
解法三
说明：
**/

func main() {
	fmt.Println(minNumber([]int{10, 2}))
}
