package main

type T struct {
	n int
}

func getT() T {
	return T{}
}

func main() {
	getT().n = 1
}

// ./test99.go:11:2: cannot assign to getT().n (value of type int)
// 无法编译
