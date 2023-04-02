package main

import (
	"fmt"
)

func main() {

	test05()
	
}

func test01()  {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[0:5]

	s2 = append(s2, 6)
	s1[3] = 30

	fmt.Println(s1)
	fmt.Println(s2)
}

func test02()  {
	s3 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s3))
	fmt.Printf("The capacity of s1: %d\n", cap(s3))
	fmt.Printf("The value of s1: %d\n", s3)
	s4 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s4))
	fmt.Printf("The capacity of s2: %d\n", cap(s4))
	fmt.Printf("The value of s2: %d\n", s4)
}

func test03()  {
	s7 := make([]int, 1024)
	fmt.Printf("The capacity of s7: %d\n", cap(s7))
	s7e1 := append(s7, make([]int, 200)...)
	fmt.Printf("s7e1: len: %d, cap: %d\n", len(s7e1), cap(s7e1))
	s7e2 := append(s7, make([]int, 400)...)
	fmt.Printf("s7e2: len: %d, cap: %d\n", len(s7e2), cap(s7e2))
	s7e3 := append(s7, make([]int, 600)...)
	fmt.Printf("s7e3: len: %d, cap: %d\n", len(s7e3), cap(s7e3))
	s9 := make([]int, 44)
	fmt.Printf("The capacity of s9: %d\n", cap(s9))
	s9a := append(s9, make([]int, 45)...)
	fmt.Printf("s9a: len: %d, cap: %d\n", len(s9a), cap(s9a))
}

func test04() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[0:3]
	// s2 := make([]int, 3)
	copy(s2, s1)
	s2 = append(s2, 40)
	s1[2] = 30

	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)

	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
}

func test05()  {
	s9 := make([]int, 100)
	fmt.Printf("len:%d, The capacity of s9: %d\n", len(s9),cap(s9))
	s9a := append(s9, make([]int, 100)...)
	fmt.Printf("s9a: len: %d, cap: %d\n", len(s9a), cap(s9a))
	fmt.Println(s9a[:223][222])
}