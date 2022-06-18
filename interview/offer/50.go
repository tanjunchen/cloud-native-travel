package main

/***
"题目：**第一个只出现一次的字符**

[第一个只出现一次的字符](https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof)

题目描述：
***/

/**
解法一
说明：
**/
func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	res := make([]int, 26)
	for _, v := range s {
		res[v-'a']++
	}
	for i := 0; i < len(s); i++ {
		if res[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

/**
解法二
说明：有序的 map
**/
