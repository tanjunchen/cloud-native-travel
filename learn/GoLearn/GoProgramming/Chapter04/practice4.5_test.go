package Chapter03

import (
	"fmt"
	"testing"
)

func removeMultiple(a *[]string) {
	A := *a
	l := len(A)
	for i := 0; i < len(A)-1; i++ {
		current := A[i]
		next := A[i+1]
		if current == next {
			A = append(A[:i], A[i+1:]...)
			l--
		}
	}
	*a = A
}
func Test045(t *testing.T) {
	b := []string{"s", "a", "a", "s", "d", "z", "a", "z", "v", "w", "w", "a", "a"}
	removeMultiple(&b)
	fmt.Println(b)
}
