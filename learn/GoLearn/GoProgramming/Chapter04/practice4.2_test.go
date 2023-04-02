package Chapter03

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"os"
	"strconv"
	"testing"
)

func encrypt(s string, b int64) string {
	var h hash.Hash
	switch b {
	case 384:
		h = sha512.New384()
	case 512:
		h = sha512.New()
	default:
		h = sha256.New()
	}
	h.Write([]byte(s))
	return string(h.Sum(nil))
}

// go test -v -run Test042 practice4.2_test.go -args 384 hello
func Test042(t *testing.T) {
	bs := os.Args[4]

	s := os.Args[5]

	b, _ := strconv.ParseInt(bs, 10, 64)

	fmt.Println(encrypt(s, b))

}


