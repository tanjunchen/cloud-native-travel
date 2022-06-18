package main

import "fmt"

func main() {
	s := []int{6, 7, 8, 10, 4, 6, 99}
	res := mergeSort(s)
	fmt.Println(res)
}

func mergeSort(s []int) []int {
	if len(s) == 1 {
		return s
	}
	m := len(s) / 2
	l := mergeSort(s[:m])
	r := mergeSort(s[m:])
	return merge(l, r)
}

func merge(l []int, r []int) []int {
	lLen, rLen := len(l), len(r)
	rs := make([]int, 0)
	lStart := 0
	rStart := 0
	for lStart < lLen && rStart < rLen {
		if l[lStart] > r[rStart] {
			rs = append(rs, r[rStart])
			rStart++
		} else {
			rs = append(rs, l[lStart])
			lStart++
		}
	}
	if lStart < lLen {
		rs = append(rs, l[lStart:]...)
	}
	if rStart < rLen {
		rs = append(rs, r[rStart:]...)
	}
	return rs
}
