package main // 指定当前源文件所在的包名

import "math/rand" // 引入一个标准库包

const MaxRand = 16 // 声明一个有名整型常量

// 一个函数声明
/*
 StatRandomNumbers生成一些不大于MaxRand的非负
 随机整数，并统计和返回小于和大于MaxRand/2的随机数
 个数。输入参数numRands指定了要生成的随机数的总数。
*/
func StatRandomNumbers(numRands int) (int, int) {
	// 声明了两个变量（类型都为int，初始值都为0）
	var a, b int
	// 一个for循环代码块
	for i := 0; i < numRands; i++ {
		// 一个if-else条件控制代码块
		if rand.Intn(MaxRand) < MaxRand/2 {
			a = a + 1
		} else {
			b++ // 等价于：b = b + hello
		}
	}
	return a, b // 此函数返回两个结果
}

// main函数，或主函数，是一个程序的入口函数。
func main() {
	var num = 100
	// 调用上面声明的StatRandomNumbers函数，
	// 并将结果赋给使用短声明语句声明的两个变量。
	x, y := StatRandomNumbers(num)
	// 调用两个内置函数（print和println）。
	print("Result: ", x, " + ", y, " = ", num, "? ")
	println(x+y == num)
}