package Chapter01

import (
	"strings"
	"testing"
)

func TestStart059(t *testing.T) {

}

func Expand(s string, f func(string) string) string {
	ret := strings.Replace(s, "$foo", f("foo"), 1024)
	return ret
}
