package main

import (
	"fmt"
	"strings"
)

func calculate(s string) int {
	arr := []byte(strings.ReplaceAll(s, " ", ""))
	var opStk []byte
	numStk := make([]int, 1)
	open := false

	for i := 0; i < len(arr); i++ {
		ln := len(numStk)
		lp := len(opStk)
		if arr[i] == byte('(') {
			open = true
		} else if arr[i] == byte('+') || arr[i] == byte('-') {
			opStk = append(opStk, arr[i])
		} else if arr[i] == byte(')') {
			numStk[ln-2] = cal(numStk[ln-2], numStk[ln-1], opStk)

			numStk = numStk[:ln-1]
			if len(opStk) > 0 {
				opStk = opStk[:lp-1]
			}
		} else {
			op2 := 0
			for len(arr) > i && arr[i] >= byte('0') && arr[i] <= byte('9') {
				op2 = op2*10 + int(arr[i]-'0')
				i++
			}
			i--
			if open {
				numStk = append(numStk, op2)
				open = false
			} else {
				numStk[ln-1] = cal(numStk[ln-1], op2, opStk)
				if len(opStk) > 0 {
					opStk = opStk[:lp-1]
				}
			}
		}
	}
	return numStk[0]
}

func cal(a, b int, op []byte) int {
	if len(op) == 0 || op[len(op)-1] == byte('+') {
		return a + b
	}
	return a - b
}

func main() {
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
