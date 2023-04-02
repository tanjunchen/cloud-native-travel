package Chapter01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func Start() {
	counts := make(map[string]int)
	files := os.Args[4:]

	if len(files) == 0 {
		log.Fatal("Input filename please")
	}

	for _, filename := range files {
		func(filename string, counts map[string]int) {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				return
			}
			countLines(f, filename, counts)
		}(filename, counts)

		for context, n := range counts {
			if n > 1 {
				filename, content := getNameAndContent(context)
				fmt.Printf("filename %s: content %q has duplicate count - %d\n", filename, content, n)
			}
		}
	}
}

func getNameAndContent(context string) (filename string, content string) {
	str := strings.Split(context, "|")
	return str[0], str[1]
}

func countLines(f *os.File, filename string, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[filename+"|"+input.Text()]++
	}
}

func TestStart(t *testing.T) {
	Start()
}
