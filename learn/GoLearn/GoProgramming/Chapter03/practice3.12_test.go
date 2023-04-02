package Chapter03

import (
	"fmt"
	"strings"
	"testing"
)

func same(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, v := range a {
		if strings.Count(a, string(v)) != strings.Count(b, string(v)) {
			return false
		}
	}
	return true
}

func Test0312(t *testing.T) {
	s1 := "bdfkajbdfkabdfj"
	s2 := "asdfaghajkl;qwertyuiop["
	s3 := ";lkjhgfdsa[aapoiuytrewq"
	s4 := ";lkjhgfdsa[poiu12rewqaa"
	fmt.Println(same(s1, s2))
	fmt.Println(same(s1, s3))
	fmt.Println(same(s2, s3))
	fmt.Println(same(s2, s4))
	fmt.Println("=================")
	fmt.Println(Isomerism(s1, s2))
	fmt.Println(Isomerism(s1, s3))
	fmt.Println(Isomerism(s2, s3))
	fmt.Println(Isomerism(s2, s4))
}

func Isomerism(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1 := make(map[string]int)
	m2 := make(map[string]int)
	for i := 0; i < len(s1); i++ {
		m1[string(s1[i])]++
		m2[string(s2[i])]++
	}
	if len(m1) != len(m2) {
		return false
	}

	for k, _ := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}

	return true
}
