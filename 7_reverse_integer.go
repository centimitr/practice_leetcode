package practice_leetcode

import "math"

func reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		// after /10, the num is an integer
		if rev > math.MaxInt32/10 || rev == math.MaxInt32/10 && pop == 7 ||
			rev < math.MinInt32/10 || rev == math.MinInt32/10 && pop < -8 {
			return 0
		}
		rev = rev*10 + pop
	}
	if x < 0 {
		rev = -rev
	}
	return rev
}
