package main

import "fmt"

/***
"题目：**131. 分割回文串**

[分割回文串](https://leetcode-cn.com/problems/palindrome-partitioning/)

给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。
*/

func partition(s string) [][]string {
	var res [][]string
	mem := make([][]int, len(s))
	for i := range mem {
		mem[i] = make([]int, len(s))
	}
	dfs([]string{}, &res, 0, s, mem)
	return res
}

func dfs(temp []string, res *[][]string, start int, s string, mem [][]int) {
	if start == len(s) {
		t := make([]string, len(temp))
		copy(t, temp)
		*res = append(*res, t)
		return
	}
	for i := start; i < len(s); i++ {
		if mem[start][i] == 2 {
			continue
		}
		if mem[start][i] == 1 || isPalindrome(s, start, i, mem) {
			temp = append(temp, s[start:i+1])
			dfs(temp, res, i+1, s, mem)
			temp = temp[:len(temp)-1]
		}
	}
}

func isPalindrome(s string, start, end int, mem [][]int) bool {
	for start < end {
		if s[start] != s[end] {
			mem[start][end] = 2
			return false
		}
		start++
		end--
	}
	mem[start][end] = 1
	return true
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
}
