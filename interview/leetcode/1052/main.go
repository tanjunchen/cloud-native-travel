package main

import "fmt"

func main() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	X := 3
	fmt.Println(maxSatisfied(customers, grumpy, X))
}

func maxSatisfied(customers []int, grumpy []int, X int) int {
	bonus := 0
	regular := 0
	for i := 0; i < X-1; i++ {
		if grumpy[i] == 1 {
			bonus += customers[i]
		} else {
			regular += customers[i]
		}
	}
	maxBonus := bonus
	for i := X - 1; i < len(customers); i++ {
		if grumpy[i] == 1 {
			bonus += customers[i]
			if bonus > maxBonus {
				maxBonus = bonus
			}
		} else {
			regular += customers[i]
		}
		if grumpy[i+1-X] == 1 {
			bonus -= customers[i+1-X]
		}
	}
	return regular + maxBonus
}
