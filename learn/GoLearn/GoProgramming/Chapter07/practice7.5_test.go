package Chapter01

import (
	"fmt"
	"io"
	"os"
	"testing"
)

type LimitedReader struct {
	Reader  io.Reader
	Limit   int
	current int
}

func (r *LimitedReader) Read(b []byte) (int, error) {
	if r.current >= r.Limit {
		return 0, io.EOF
	}

	if r.current+len(b) > r.Limit {
		b = b[:r.Limit-r.current]
	}

	n, err := r.Reader.Read(b)
	if err != nil {
		return n, err
	}
	r.current += n
	return n, nil
}

func LimitReader(r io.Reader, limit int) io.Reader {
	lr := LimitedReader{
		Reader: r,
		Limit:  limit,
	}
	return &lr
}

func Test075(t *testing.T) {
	file, err := os.Open("limit.txt") // 1234567890
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lr := LimitReader(file, 5)
	buf := make([]byte, 10)
	n, err := lr.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, buf) // 5 [49 50 51 52 53 0 0 0 0 0]
}
