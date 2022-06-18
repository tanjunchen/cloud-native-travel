package main

/***
"题目：**395. 至少有 K 个重复字符的最长子串**

[至少有 K 个重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/)

给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。

***/

/**
解法一
说明：递归循环
**/

func longestSubstring(s string, k int) int {
	return helper(0, len(s)-1, k, s)
}

func helper(start int, end int, k int, s string) int {
	if end-start+1 < k {
		return 0
	}
	res := make(map[byte]int, end-start+1)
	for i := start; i <= end; i++ {
		res[s[i]]++
	}
	for end-start+1 >= k && res[s[start]] < k {
		start++
	}
	for end-start+1 >= k && res[s[end]] < k {
		end--
	}
	if end-start+1 < k {
		return 0
	}
	for i := start; i <= end; i++ {
		if res[s[i]] < k {
			return max(helper(start, i-1, k, s), helper(i+1, end, k, s))
		}
	}
	return end - start + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
