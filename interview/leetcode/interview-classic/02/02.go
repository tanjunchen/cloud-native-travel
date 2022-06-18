package main

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[rune]int)
	for _, i := range s1 {
		m[i]++
	}
	for _, i := range s2 {
		m[i]--
	}
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}

func main() {
}
