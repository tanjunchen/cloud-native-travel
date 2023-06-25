package main

// 有关 channel 说法正确的是：
// A. 向已关闭的通道发送数据会引发 panic；
// B. 从已关闭的缓冲通道接收数据，返回已缓冲数据或者零值；
// C. 无论接收还是接收，nil 通道都会阻塞；

func main() {
	x := map[string]string{"one": "a", "two": "", "three": "c"}
	if v := x["two"]; v == "" {
		fmt.Println("no entry")
	}
}

// 检查 map 是否含有某一元素，直接判断元素的值并不是一种合适的方式。
// 最可靠的操作是使用访问 map 时返回的第二个值。
