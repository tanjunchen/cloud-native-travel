package main

import (
	"fmt"
)

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

// 在使用 fmt 包中的打印方法时，如果类型实现了这个接口，会直接调用。产生循环调用。
func main() {
	p := &People{}
	p.String()
}
