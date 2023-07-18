package main

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
}

func flipAndInvertImage(A [][]int) [][]int {
	if len(A) == 0 || A == nil {
		return A
	}
	m, n := len(A), len(A[0])
	for i := 0; i < m; i++ {
		for j, k := 0, n-1; j <= k; j, k = j+1, k-1 {
			A[i][k], A[i][j] = 1^A[i][j], 1^A[i][k]
		}
	}
	return A
}
