package main

func oneEditAway(first string, second string) bool {
	if first == second {
		return true
	}
	if abs(len(first), len(second)) {
		return false
	}
	p1, p2 := 0, 0
	flag := false
	for p1 < len(first) || p2 < len(second) {
		if p1 < len(first) && p2 < len(second) && first[p1] == second[p2] {
			p1++
			p2++
			continue
		}
		if !flag {
			if len(first)-p1 > len(second)-p2 {
				p1++
			} else if len(first)-p1 < len(second)-p2 {
				p2++
			} else {
				p1++
				p2++
			}
			flag = true
		} else {
			return false
		}
	}
	if p1 == len(first) && p2 == len(second) {
		return true
	}
	return false
}

func abs(a, b int) bool {
	if a > b {
		return a-b > 1
	}
	return b-a > 1
}
