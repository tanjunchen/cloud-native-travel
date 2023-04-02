package Chapter03

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"testing"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Test043(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)
	b := []int{0, 1, 2, 3, 4, 5, 6, 7}
	reverse2(&b)
	fmt.Println(b)
}

func reverse2(arr *[]int) {
	a := *arr
	l := len(a)
	for i := 0; i < l/2; i++ {
		a[i], a[l-i-1] = a[l-i-1], a[i]
	}
}

func analysisParams() {

	method := flag.String("method", "sha256", "select hash method(sha256,sha384,sha512)")

	text := flag.String("text", "", "input the string your want to hash")

	flag.Parse()

	switch *method {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(*text)))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(*text)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(*text)))
	default:
		panic("not support hash method")
	}

}

// go test -v -run Test0432 practice4.3_test.go  -args sha512
func Test0432(t *testing.T) {
	analysisParams()
}
