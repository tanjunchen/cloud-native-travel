package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compressString(S string) string {
	if len(S) == 0 {
		return ""
	}
	var str strings.Builder
	temp := S[0]
	count := 1
	for i := 1; i < len(S); i++ {
		if temp == S[i] {
			count++
		} else {
			str.WriteByte(temp)
			str.WriteString(strconv.Itoa(count))
			count = 1
			temp = S[i]
		}
	}
	str.WriteByte(temp)
	str.WriteString(strconv.Itoa(count))
	if str.Len() >= len(S) {
		return S
	}
	return str.String()
}

func main() {
	fmt.Println(compressString("aabcccccaaa"))
	fmt.Println(compressString("abbccd"))
}
