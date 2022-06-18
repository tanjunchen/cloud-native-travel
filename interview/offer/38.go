package main

import (
	"fmt"
	"sort"
)

/***
"题目：**字符串的排列**

[字符串的排列](https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof)

题目描述：输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
***/

/**
解法一
说明：
**/
func permutation(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	bs := []byte(s)
	var temp []int
	for i := 0; i < len(bs); i++ {
		temp = append(temp, int(bs[i]))
	}
	sort.Ints(temp)
	for i := 0; i < len(temp); i++ {
		bs[i] = byte(temp[i])
	}
	s = string(bs)
	var res []string
	var list []byte
	visited := make(map[int]bool)
	for i := 0; i < len(s); i++ {
		visited[i] = false
	}
	backTrack(&res, &list, s, visited)
	return res
}

func backTrack(res *[]string, list *[]byte, s string, visited map[int]bool) {
	if len(s) == len(*list) {
		*res = append(*res, string(*list))
	}
	for i := 0; i < len(s); i++ {
		if visited[i] == true {
			continue
		}
		if i >= 1 && visited[i-1] == false && s[i] == s[i-1] {
			continue
		}
		visited[i] = true
		*list = append(*list, s[i])
		backTrack(res, list, s, visited)
		visited[i] = false
		*list = (*list)[:len(*list)-1]
	}
}

func main() {
	fmt.Println(permutation("abc"))
}
