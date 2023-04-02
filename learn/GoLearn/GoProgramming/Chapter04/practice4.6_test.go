package Chapter03

import (
	"fmt"
	"testing"
	"unicode"
)

func removeEmpty(a *string) {
	A := *a
	begin := 0
	ss := string(A[0])
	for _, s := range A {
		current := rune(ss[begin])
		if unicode.IsSpace(current) && unicode.IsSpace(s) {
			continue
		}
		ss += string(s)
		begin++
	}
	*a = ss[1:]
}

func Test046(t *testing.T) {
	a := "asdf    sadf  sd f a   sfd    a"
	removeEmpty(&a)
	fmt.Println(a)
	b := "aa   aaa aaa"
	bb := DeleteEmpty([]rune(b))
	fmt.Println(string(bb))
}

func DeleteEmpty(s []rune) []rune {
	position := 0
	tag := unicode.IsSpace(s[0])
	for i := 1; i < len(s); i++ {
		if unicode.IsSpace(s[i]) {
			if tag != true {
				position += 1
				s[position] = s[i]
				tag = true
			}
		} else {
			tag = false
			position += 1
			s[position] = s[i]
		}
	}
	return s[:position+1]
}
