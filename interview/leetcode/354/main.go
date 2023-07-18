package main

import (
	"fmt"
	"sort"
)

/***
"题目：**354 俄罗斯套娃信封问题**

[俄罗斯套娃信封问题](https://leetcode-cn.com/problems/russian-doll-envelopes/)

给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。

当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

注意：不允许旋转信封。

***/
/**
基于二分查找的动态规划   经典的「最长严格递增子序列」
*/
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	var f []int
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}

func main() {
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
	fmt.Println(maxEnvelopes([][]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}}))
}
