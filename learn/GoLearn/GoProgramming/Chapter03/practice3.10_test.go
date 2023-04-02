package Chapter03

import (
	"bytes"
	"fmt"
	"testing"
)

func comma(str string) string {
	n := len(str)
	if n <= 3 {
		return str
	}
	return comma(str[:n-3]) + "," + str[n-3:]
}

func Test0310(t *testing.T) {
	s := "12345"
	fmt.Println(comma(s))
	fmt.Println(comma2(s))
}

func comma2(s string) string {
	var r bytes.Buffer

	l := len(s)
	mod := l % 3
	if mod > 0 {
		r.Write([]byte(s[:mod] + ","))
	}
	for mod+3 < l {
		r.Write([]byte(s[mod:mod+3] + ","))
		mod += 3
	}
	if mod+3 == l {
		r.Write([]byte(s[mod : mod+3]))
	}
	return r.String()
}
