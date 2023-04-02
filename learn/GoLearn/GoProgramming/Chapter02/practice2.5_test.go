package Chapter02

import (
	"testing"
)

func PopCount025(x uint64) int {
	ret := 0
	for x != 0 {
		x = (x - 1) & x
		ret += 1
	}
	return ret
}

func BenchmarkPopCount025(b *testing.B) {
	// BenchmarkPopCount025-12    	40105879	        29.3 ns/op
	for i := 0; i < b.N; i++ {
		PopCount025(2<<63 - 1)
	}
}
