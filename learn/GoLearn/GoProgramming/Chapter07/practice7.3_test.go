package Chapter01

import (
	"bytes"
	"fmt"
)

type tree struct {
	Val   int
	Left  *tree
	Right *tree
}

func (t *tree) String() string {
	var deque []*tree
	var result []int
	deque = append(deque, t)
	for len(deque) > 0 {
		current := deque[0]
		deque = deque[1:]
		if current.Left != nil {
			deque = append(deque, current.Left)
		}
		if current.Right != nil {
			deque = append(deque, current.Right)
		}
		result = append(result, current.Val)
	}

	var buf bytes.Buffer
	buf.Write([]byte("{"))
	for i, v := range result {
		if i == len(result)-1 {
			buf.Write([]byte(fmt.Sprintf("%d", v)))
		} else {
			buf.Write([]byte(fmt.Sprintf("%d, ", v)))
		}
	}
	buf.Write([]byte("}"))
	return buf.String()
}
