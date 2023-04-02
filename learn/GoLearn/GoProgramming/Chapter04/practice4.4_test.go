package Chapter03

import (
	"fmt"
	"testing"
)

func rotate(arr []int, pos int) []int {
	r := arr[pos:]
	for i := pos - 1; i >= 0; i-- {
		r = append(r, arr[i])
	}
	return r
}
func Test044(t *testing.T) {
	b := []int{0, 1, 2, 3, 4, 5, 6, 7}
	r := rotate(b, 3)
	fmt.Println(r)
}
