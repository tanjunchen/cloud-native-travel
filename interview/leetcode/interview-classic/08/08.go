package main

import "fmt"

func setZeroes(matrix [][]int) [][]int {
	var flags [][]int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				flags = append(flags, []int{i, j})
			}
		}
	}
	for i := 0; i < len(flags); i++ {
		for k := 0; k < len(matrix[flags[i][0]]); k++ {
			matrix[flags[i][0]][k] = 0
		}

		for l := 0; l < len(matrix); l++ {
			matrix[l][flags[i][1]] = 0
		}
	}
	return matrix
}

func main() {
	fmt.Println(setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
}
