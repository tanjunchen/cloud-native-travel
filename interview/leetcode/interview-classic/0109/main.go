package main

import (
	"fmt"
	"strings"
)

func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s2+s2, s1)
}

func isFlipedString2(s1 string, s2 string) bool {

	len1 := len(s1)
	len2 := len(s2)
	if len1 != len2 {
		return false
	}
	if len1 == 0 {
		return true
	}

	j := 0
	for i := 0; i < 2*len1; i++ {
		if j == len2 {
			return true
		}
		if s1[i%len1] == s2[j] {
			j++
		} else {
			j = 0
		}
	}
	return false
}

func main() {
	fmt.Println(isFlipedString("waterbottle", "erbottlewat"))
	fmt.Println(isFlipedString2("waterbottle", "erbottlewat"))
	fmt.Println(isFlipedString2("aa", "aba"))
}
