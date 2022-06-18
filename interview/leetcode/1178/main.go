package main

import "fmt"

/***
"题目：**1178.猜字谜**

[猜字谜](https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle/)

外国友人仿照中国字谜设计了一个英文版猜字谜小游戏，请你来猜猜看吧。

字谜的迷面 puzzle 按字符串形式给出，如果一个单词 word 符合下面两个条件，那么它就可以算作谜底：

单词 word 中包含谜面 puzzle 的第一个字母。
单词 word 中的每一个字母都可以在谜面 puzzle 中找到。
例如，如果字谜的谜面是 "abcdefg"，那么可以作为谜底的单词有 "faced", "cabbage", 和 "baggage"；而 "beefed"（不含字母 "a"）
以及 "based"（其中的 "s" 没有出现在谜面中）都不能作为谜底。
返回一个答案数组 answer，数组中的每个元素 answer[i] 是在给出的单词列表 words 中可以作为字谜迷面 puzzles[i] 所对应的谜底的单词数目。
***/

/**
解法一
words = ["aaaa","asas","able","ability","actt","actor","access"],
puzzles = ["aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"]
**/

var res map[int]int

func findNumOfValidWords(words []string, puzzles []string) []int {
	res = map[int]int{}
	for _, w := range words {
		res[bset(w)] += 1
	}
	ans := make([]int, len(puzzles))
	for i, p := range puzzles {
		ans[i] = f(p)
	}
	return ans
}

func bset(w string) int {
	var v int
	for _, c := range w {
		v |= 1 << (c - 97)
	}
	return v
}

func f(p string) (sum int) {
	cbt := []string{p[0:1]}
	n := len(p)
	for i := 1; i < n; i++ {
		var tmp []string
		for _, s := range cbt {
			tmp = append(tmp, s+p[i:i+1])
		}
		cbt = append(cbt, tmp...)
	}
	for _, s := range cbt {
		if v, ok := res[bset(s)]; ok {
			sum += v
		}
	}
	return
}

func main() {
	//fmt.Println(findNumOfValidWords([]string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"},[]string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}))
	test("ability")
}

func test(p string) {
	cbt := []string{p[0:1]}
	n := len(p)
	for i := 1; i < n; i++ {
		var tmp []string
		for _, s := range cbt {
			tmp = append(tmp, s+p[i:i+1])
		}
		cbt = append(cbt, tmp...)
	}
	fmt.Println(cbt)
}
