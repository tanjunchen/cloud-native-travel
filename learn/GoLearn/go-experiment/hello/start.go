package main

import (
	"fmt"
	"time"
)

func testswitch() {
	i := 2
	fmt.Println("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("one")
	case 3:
		fmt.Println("one")
	default:
		fmt.Println("cccc")
	}

	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday:
		fmt.Println("it is the weekend")
	default:
		fmt.Println("it is the weekday")
	}
}

func main() {
	testswitch()
}
