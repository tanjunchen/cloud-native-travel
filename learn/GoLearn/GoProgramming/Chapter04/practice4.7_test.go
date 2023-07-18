package Chapter03

import (
	"fmt"
	"testing"
)

func reverse47(arr []byte) string {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
	return string(arr)
}
func Test047(t *testing.T) {
	a := "hello world!"
	fmt.Println(reverse47([]byte(a)))
	Reverse()
}

func Reverse() {
	a := []byte("我要搞Golang")
	b := []rune(string(a))
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	fmt.Println(string(b))
}
