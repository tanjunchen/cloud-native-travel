package Chapter01

import (
	"fmt"
	"io"
	"testing"
)

func Test074(t *testing.T) {
	str := "Hello World"
	sr := NewReader(str)
	data := make([]byte, 10)
	n, err := sr.Read(data)
	for err == nil {
		fmt.Println(n, string(data[0:n]))
		n, err = sr.Read(data)
	}
}

type StringReader struct {
	data string
	n    int
}

func (sr *StringReader) Read(b []byte) (int, error) {
	data := []byte(sr.data)
	if sr.n >= len(data) {
		return 0, io.EOF
	}
	n := 0
	if len(b) >= len(data) {
		n = copy(b, data)
		sr.n = sr.n + n
		return n, nil
	}
	data = data[sr.n:]
	n = copy(b, data)
	sr.n = sr.n + n
	return n, nil
}

func NewReader(in string) *StringReader {
	sr := new(StringReader)
	sr.data = in
	return sr
}

