package main

/**
给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool，
代表两串是否重新排列后可相同。
保证两串的⻓度都小于等于5000。
*/

func main() {
	s1 := "abcdefghijklmnopqrstuvwxyz"
	s2 := "sssssssssss"
}

func test(s1, s2 string) bool {
	s1rune := []rune(s1)
	s2rune := []rune(s2)
	l1 := len(s1rune)
	l2 := len(s2rune)
	if l1 > 5000 || l2 > 5000 || l1 != l2 {
		return false
	}

}
