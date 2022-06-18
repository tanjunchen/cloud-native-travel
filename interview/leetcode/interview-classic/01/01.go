package main

import "fmt"

func isUnique(astr string) bool {
	res := make(map[byte]bool)
	for _, i := range astr {
		if _, ok := res[byte(i)]; ok {
			return false
		}
		res[byte(i)] = true
	}
	return true
}

func main() {
	fmt.Println(isUnique("leetcode"))
	fmt.Println(isUnique("abc"))
}
