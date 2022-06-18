package main

import "fmt"

func removeDuplicates(S string) string {
	var stacks []byte
	for i := range S {
		if len(stacks) > 0 && stacks[len(stacks)-1] == S[i] {
			stacks = stacks[:len(stacks)-1]
		} else {
			stacks = append(stacks, S[i])
		}
	}
	return string(stacks)
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}
