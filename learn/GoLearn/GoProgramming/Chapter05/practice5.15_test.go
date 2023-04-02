package Chapter01

import (
	"math"
	"testing"
)

func TestStart015(t *testing.T) {

}

func Max(in ...int) int {
	if len(in) == 0 {
		panic("At least one element")
	}
	ret := math.MinInt64
	for _, v := range in {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func Min(in ...int) int {
	if len(in) == 0 {
		panic("At least one element")
	}

	ret := math.MaxInt64
	for _, v := range in {
		if v < ret {
			ret = v
		}
	}
	return ret
}
