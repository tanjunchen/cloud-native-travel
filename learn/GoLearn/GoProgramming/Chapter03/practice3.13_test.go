package Chapter03

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func comma6(s string) {
	var r bytes.Buffer
	start := ""

	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		start = string(s[0])
		s = s[1:]
	}

	end := ""
	if strings.Contains(s, ".") {
		ss := strings.Split(s, ".")
		s, end = ss[0], "."+ss[1]
	}

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
	fmt.Println(strings.Join([]string{start, r.String(), end}, ""))
}

func Test0313(t *testing.T) {
	comma6("12345")
	comma6("12345.hello")
	comma6("-12345.hello")
}
