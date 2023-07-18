package main

/**
哈希值
*/
func canPermutePalindrome(s string) bool {
	res := make(map[rune]int)
	for _, i := range s {
		res[i] = res[i] + 1
	}
	count := 0
	for _, v := range res {
		if v%2 != 0 {
			count++
		}
	}
	return count == 0 || count == 1
}
