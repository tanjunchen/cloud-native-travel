package Chapter03

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func start049() {
	wordCount := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordCount[input.Text()]++
	}
	fmt.Printf("\nWorld\tCount\n")
	for key, value := range wordCount {
		fmt.Printf("%v\t%d\n", key, value)
	}
}
// go test -v -run Test049 practice4.9_test.go  -args A A B C
func Test049(t *testing.T) {
	start049()
}
