package Chapter03

import (
	"fmt"
	"math"
	"testing"
)

func Test031(t *testing.T) {
	var z float64
	if 1/z > math.MaxFloat64 {
		fmt.Println(math.MaxFloat64)
	}
}
