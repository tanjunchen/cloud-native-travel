package Chapter02

import (
	"testing"
)

func PopCount024(x uint64) int {
	ret := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			ret += 1
		}
		x = x >> 1
	}
	return ret
}

// BenchmarkPopCount024-12    	32518738	        39.0 ns/op
func BenchmarkPopCount024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount024(2<<63 - 1)
	}
}
