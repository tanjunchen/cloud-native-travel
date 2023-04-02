package Chapter01

import (
	"bytes"
	"strings"
	"testing"
)

func TestStart016(t *testing.T) {

}

func Join(in ...string) string {
	return strings.Join(in, "")
}

func Join2(in ...string) string {
	var buf bytes.Buffer
	for _, v := range in {
		buf.Write([]byte(v))
	}
	return buf.String()
}
