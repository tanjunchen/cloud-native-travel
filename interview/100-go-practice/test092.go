package main

type X struct{}

func (x *X) test() {
	println(x)
}

func main() {
	// 该方法无问题
	// var a *X
	// a.test()
	// 下述方法有问题
	// X{}.test()
	var x = X{}
	x.test()
}
