package Chapter01

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func Test071(t *testing.T) {
	var wc WordsCounter
	wc.Write([]byte("Hello Worlds Test Me"))
	fmt.Println(wc) // 4
	wc.Write([]byte("append something to the end"))
	fmt.Println(wc) // 9

	var lc LinesCounter
	fmt.Fprintf(&lc, "%s\n%s\n%s\n", "Hello World", "Second Line", "Third Line")
	fmt.Println(lc) // 3
	fmt.Fprintf(&lc, "%s\n%s\n%s", "第4行", "第5行", "")
	fmt.Println(lc) // 5
}

type WordsCounter int

func (c *WordsCounter) Write(content []byte) (int, error) {
	for start := 0; start < len(content); {
		advance, _, err := bufio.ScanWords(content[start:], true)
		if err != nil {
			return 0, err
		}
		start += advance
		*c++
	}
	return int(*c), nil
}

type LinesCounter int

func (c *LinesCounter) Write(content []byte) (int, error) {
	for start := 0; start < len(content); {
		advance, _, err := bufio.ScanLines(content[start:], true)
		if err != nil {
			return 0, err
		}
		start += advance
		*c++
	}
	return int(*c), nil
}

type CountWriter struct {
	Writer io.Writer
	Count  int
}

func (cw *CountWriter) Write(content []byte) (int, error) {
	n, err := cw.Writer.Write(content)
	if err != nil {
		return n, err
	}
	cw.Count += n
	return n, nil
}

func CountingWriter(writer io.Writer) (io.Writer, *int) {
	cw := CountWriter{
		Writer: writer,
	}
	return &cw, &(cw.Count)
}

func Test072(t *testing.T) {
	cw, counter := CountingWriter(os.Stdout)
	fmt.Fprintf(cw, "%s", "Print something to the screen...")
	fmt.Println(*counter)

	cw.Write([]byte("AAA"))
	fmt.Println(*counter)
}