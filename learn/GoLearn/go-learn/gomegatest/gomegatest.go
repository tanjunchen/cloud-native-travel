package gomegatest

import "fmt"

func StartGomega() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, byte('a')))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	length := len(letters)

	if letters[0] > target || letters[length-1] <= target {
		return letters[0]
	}

	l, r := 0, length-1
	for l <= r {
		m := l + (r-l)/2
		if letters[m] <= target {
			if m == length-1 {
				return letters[m]
			}
			if letters[m+1] > target {
				return letters[m+1]
			} else {
				l = m + 1
			}
		}else{
			r = m - 1
		}
	}
	return ' '
}
