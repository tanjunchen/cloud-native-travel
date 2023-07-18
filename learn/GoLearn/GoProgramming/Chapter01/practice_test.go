package Chapter01

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestTest111(t *testing.T) {
	fmt.Println(os.Args[0])
}

func TestTest112(t *testing.T) {
	for k, v := range os.Args {
		fmt.Printf("%d\t%s\n", k, v)
	}
}

func TestTest113(t *testing.T) {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		fmt.Println(i, args[i])
	}
}

func BenchmarkStringJoin1(b *testing.B) {
	var a [100000]string
	for i := 0; i < 100000; i++ {
		a[i] = "A"
	}
	for i := 0; i < b.N; i++ {
		var tmp string
		for _, str := range a {
			tmp += str
		}
	}
}

func BenchmarkStringJoin2(b *testing.B) {
	var a [100000]string
	for i := 0; i < 100000; i++ {
		a[i] = "A"
	}

	for i := 0; i < b.N; i++ {
		_ = strings.Join(a[:], "")
	}
}

func TestCal(t *testing.T) {
	printArgs1()
	printArgs2()
}
func TimeConsuming(tag string) func() {
	now := time.Now().UnixNano()
	return func() {
		after := time.Now().UnixNano()
		fmt.Printf("%q time cost %d ns\n", tag, after-now)
	}
}

func printArgs1() {
	defer TimeConsuming("printArgs1")()
	fmt.Println(os.Args[1:])
}

func printArgs2() {
	defer TimeConsuming("printArgs2")()
	args := strings.Join(os.Args[1:], " ")
	fmt.Println(args)
}
