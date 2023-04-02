package Chapter01

import (
	"fmt"
	"testing"
)

func TestStart019(t *testing.T) {
	fmt.Println(NoReturn())
}

func NoReturn() (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = 1
		} else {
			r = 2
		}
	}()
	panic(1)
}
