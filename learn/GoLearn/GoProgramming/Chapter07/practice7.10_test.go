package Chapter01

import (
	"fmt"
	"sort"
	"testing"
)

func IsPalindrome(s sort.Interface) bool {
	if s.Len() == 0 {
		return false
	}

	i, j := 0, s.Len()-1
	for i < j {
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func Test079(t *testing.T) {
	start001()
}

func start001() {
	a := []int{1, 2, 3, 2, 1}
	fmt.Println(IsPalindrome(sort.IntSlice(a))) // true
	a = []int{2, 1, 3, 4, 5}
	fmt.Println(IsPalindrome(sort.IntSlice(a))) //false
	a = []int{1}
	fmt.Println(IsPalindrome(sort.IntSlice(a))) // true
	a = []int{}
	fmt.Println(IsPalindrome(sort.IntSlice(a))) // false
}
