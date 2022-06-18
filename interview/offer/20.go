package main

import (
	"regexp"
	"strings"
)

/***
"题目：**表示数值的字符串**

[表示数值的字符串](https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof)

题目描述：请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。

***/

/**
解法一
说明：
**/
// 确定有限状态自动机
type State int
type CharType int

const (
	STATE_INITIAL State = iota
	STATE_INT_SIGN
	STATE_INTEGER
	STATE_POINT
	STATE_POINT_WITHOUT_INT
	STATE_FRACTION
	STATE_EXP
	STATE_EXP_SIGN
	STATE_EXP_NUMBER
	STATE_END
)

const (
	CHAR_NUMBER CharType = iota
	CHAR_EXP
	CHAR_POINT
	CHAR_SIGN
	CHAR_SPACE
	CHAR_ILLEGAL
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER
	case 'e', 'E':
		return CHAR_EXP
	case '.':
		return CHAR_POINT
	case '+', '-':
		return CHAR_SIGN
	case ' ':
		return CHAR_SPACE
	default:
		return CHAR_ILLEGAL
	}
}

func isNumber(s string) bool {
	transfer := map[State]map[CharType]State{
		STATE_INITIAL: {
			CHAR_SPACE:  STATE_INITIAL,
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
			CHAR_SIGN:   STATE_INT_SIGN,
		},
		STATE_INT_SIGN: {
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
		},
		STATE_INTEGER: {
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_EXP:    STATE_EXP,
			CHAR_POINT:  STATE_POINT,
			CHAR_SPACE:  STATE_END,
		},
		STATE_POINT: {
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
			CHAR_SPACE:  STATE_END,
		},
		STATE_POINT_WITHOUT_INT: {
			CHAR_NUMBER: STATE_FRACTION,
		},
		STATE_FRACTION: {
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
			CHAR_SPACE:  STATE_END,
		},
		STATE_EXP: {
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SIGN:   STATE_EXP_SIGN,
		},
		STATE_EXP_SIGN: {
			CHAR_NUMBER: STATE_EXP_NUMBER,
		},
		STATE_EXP_NUMBER: {
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SPACE:  STATE_END,
		},
		STATE_END: {
			CHAR_SPACE: STATE_END,
		},
	}
	state := STATE_INITIAL
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = transfer[state][typ]
		}
	}
	return state == STATE_INTEGER || state == STATE_POINT || state == STATE_FRACTION || state == STATE_EXP_NUMBER || state == STATE_END
}

// 正则表达式 好复杂
func isNumber2(s string) bool {
	s = strings.TrimSpace(s)
	// res,_ := regexp.MatchString("^(([\\+\\-]?[0-9]+(\\.[0-9]*)?)|([\\+\\-]?\\.?[0-9]+))(e[\\+\\-]?[0-9]+)?$", s)
	res, _ := regexp.MatchString("^(([\\+\\-]?[0-9]+(\\.[0-9]*)?)|([\\+\\-]?\\.?[0-9]+))((e|E)[\\+\\-]?[0-9]+)?$", s)
	return res
}

func isNumber3(s string) bool {
	/**
	如果有指数记号e/E，那么指数记号的左边可以为浮点数，右边不能为浮点数，例如5.1e1是正确的，5e1.1是错误的
	如果有指数记号e/E，那么指数记号的左右两边都必须有数字出现，例如5e1是正确的，5e或e1是错误的
	如果有指数记号e/E，那么指数记号的左右两边各自最多只能出现一个正负号，且正负号必须出现在这一部分的最前面，例如+1是正确的，+-1或2-1是错误的
	最多只能出现一个小数点
	不能出现+、-、.、e、E和0~9之外的字符，如果有空格，空格必须出现在最前面和最后面
	*/
	// 我们可以使用状态机，用四个变量分别记录是否出现了指数标记、是否出现了数字、是否出现了正负号、是否出现了小数点
	// 串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。
	s = strings.Trim(s, " ")
	findE := false
	findNum := false
	findSignal := false
	findDot := false
	for _, c := range s {
		if c == '+' || c == '-' {
			if findSignal {
				return false
			}
			if findNum {
				return false
			}
			if findDot {
				return false
			}
			findSignal = true
		} else if c == 'E' || c == 'e' {
			if !findNum {
				return false
			}
			if findE {
				return false
			}
			findDot = false
			findE = true
			findNum = false
			findSignal = false
		} else if c == '.' {
			if findE {
				return false
			}
			if findDot {
				return false
			}
			findDot = true
		} else if c <= '9' && c >= '0' {
			findNum = true
		} else {
			return false
		}
	}
	return findNum
}

func main() {

}
