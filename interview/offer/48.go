package main

/***
"题目：**最长不含重复字符的子字符串**

[最长不含重复字符的子字符串](https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof)

题目描述：请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度
***/

/**
解法一
说明：
**/
func lengthOfLongestSubstring(s string) int {
	// 滑动窗口的思想
	l, r := 0, 0
	data := make(map[byte]int)
	res := 0
	for r < len(s) {
		data[s[r]] += 1
		r++
		for l < r && data[s[r-1]] == 2 {
			if data[s[l]] == 2 {
				data[s[l]] -= 1
				l++
				break
			} else {
				data[s[l]] -= 1
				l++
			}
		}
		if r-l > res {
			res = r - l
		}
	}
	return res
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 滑动窗口的思想
	l := 0
	data := make(map[byte]bool)
	res := 1
	data[s[0]] = true
	for r := 1; r < len(s); {
		flag := data[s[r]]
		if !flag {
			data[s[r]] = true
			r++
		} else {
			delete(data, s[l])
			l++
		}
		if r-l > res {
			res = r - l
		}
	}
	return res
}
