package main

import (
	"fmt"
	"math"
	"strings"
)

/***
"题目：**把字符串转换成整数**

[把字符串转换成整数](https://leetcode-cn.com/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof)

题目描述： 写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。

首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。

当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。

注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。

在任何情况下，若函数不能进行有效的转换时，请返回 0。

说明：

假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
***/

/**
解法一
说明：
**/
func strToInt(str string) int {
	// 根据长度判断边界
	// 去除空格
	// 判断首字母是否为非数字或者正负符号
	// 组装数字
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	flag := 1
	result := 0
	for i, v := range str {
		if i == 0 && v == '-' {
			flag = -flag
		} else if i == 0 && v == '+' {
			flag = flag
		} else if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else {
			break
		}

		if result > math.MaxInt32 {
			if flag == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	if flag == -1 {
		return -result
	}
	return result
}

/**
解法二
说明：模仿 aoti 函数
**/
func strToInt2(str string) int {
	// 根据长度判断边界，转换为 byte 数组
	// 去除空格
	// 判断首字母是否为非数字或者正负符号
	// 组装数字
	if len(str) == 0 {
		return 0
	}
	bts := []byte(str)
	start := 0
	ch := bts[0]
	for ch == ' ' && start < len(bts)-1 {
		start++
		ch = bts[start]
	}
	bts = bts[start:]
	if len(bts) == 0 || (bts[0] >= 'a' && bts[0] <= 'z') || (bts[0] >= 'A' && bts[0] <= 'Z') {
		return 0
	}
	flag := false
	if bts[0] == '-' {
		flag = true
		bts = bts[1:]
	} else if bts[0] == '+' {
		bts = bts[1:]
	}
	res := 0
	for _, v := range bts {
		if res > (1<<31)-1 {
			if flag {
				return -2147483648
			} else {
				return (1 << 31) - 1
			}
		}
		if v >= '0' && v <= '9' {
			res = res*10 + int(v-'0')
		} else {
			break
		}
	}
	if res > (1<<31)-1 {
		if flag {
			return -2147483648
		} else {
			return (1 << 31) - 1
		}
	}
	if flag {
		res = -res
	}
	return res
}

func main() {
	fmt.Println(strToInt2("    -42"))
	fmt.Println(strToInt2("42"))
	fmt.Println(strToInt2("4193 with words"))
	fmt.Println(strToInt2("words and 987"))
	fmt.Println(strToInt2("-91283472332"))
	fmt.Println(strToInt2("3.1415"))
	fmt.Println(strToInt2("2147483648"))   // 2147483647
	fmt.Println(strToInt2("-91283472332")) // -2147483648
}
