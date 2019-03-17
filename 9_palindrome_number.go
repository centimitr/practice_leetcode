package practice_leetcode

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	b := []byte(strconv.Itoa(x))
	l := len(b)
	for i := 0; i < l/2; i++ {
		if b[i] != b[l-i-1] {
			return false
		}
	}
	return true
}
