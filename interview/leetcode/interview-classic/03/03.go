package main

import (
	"fmt"
	"strings"
)

func replaceSpaces(S string, length int) string {
	S = S[:length]
	return strings.ReplaceAll(S, " ", "%20")
}

func replaceSpaces2(S string, length int) string {
	var str strings.Builder
	for i := 0; i < length; i++ {
		if string(S[i]) == " " {
			str.WriteString("%20")
		} else {
			str.WriteByte(S[i])
		}
	}
	return str.String()
}

func main() {
	fmt.Println(replaceSpaces("Mr John Smith    ", 13))
	fmt.Println(replaceSpaces("               ", 5))
	fmt.Println(replaceSpaces2("Mr John Smith    ", 13))
	fmt.Println(replaceSpaces2("               ", 5))
}
