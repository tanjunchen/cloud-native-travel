package Chapter02

import (
	"fmt"
	"testing"
	"time"
)

var pc [512]byte

func init() {
	for i := range pc {
		pc[i] = pc[ i/2] + byte(i&1)
	}
}
func TimeConsuming(tag string) func() {
	now := time.Now().UnixNano() / 1000000
	return func() {
		after := time.Now().UnixNano() / 1000000
		fmt.Printf("%q time cost %d ms\n", tag, after-now)
	}
}

func PopCount1(x uint64) int {
	defer TimeConsuming("PopCount1")()
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	defer TimeConsuming("PopCount2")()
	var n int
	for i := 0; i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}

func Test023(t *testing.T) {
	fmt.Println(PopCount1(200)) // "PopCount1" time cost 2000 ns
	fmt.Println(PopCount2(200)) // "PopCount2" time cost 0 ns
}

func PopCount3(x uint64) int {
	ret := byte(0)
	for i := uint8(0); i < 8; i++ {
		ret += pc[byte(x>>(i*8))]
	}
	return int(ret)
}

/**
书上的代码主要是预先计算,将64bit每8bit一组,8bit会有二的八次方种结果，共256种结果，将所有的结果都先计算出来(空间换时间)

预计算主要是这一步pc[i] = pc[i/2] + byte(i&hello)，注意此处的含义是除以二表示二进制数右移一位，后面表示最后一位是1还是0

代码中[256]byte换成[256]int，更好理解。可能这样写只是为了让人更加能搞明白byte和int的各种转换吧

注意这种预计算的执行效率是很高的
*/

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func BenchmarkPopCount3(b *testing.B) {
	// BenchmarkPopCount3-12    	85711224	        14.4 ns/op
	for i := 0; i < b.N; i++ {
		PopCount3(2<<63 - 1)
	}
}

func BenchmarkPopCount(b *testing.B) {
	// BenchmarkPopCount-12    	1000000000	         0.254 ns/op
	for i := 0; i < b.N; i++ {
		PopCount(2<<63 - 1)
	}
}
