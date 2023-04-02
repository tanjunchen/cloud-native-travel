package forwhile

import "fmt"

func Forwhile() {
	i := 1
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}
}

func IfElse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		
	}
}
