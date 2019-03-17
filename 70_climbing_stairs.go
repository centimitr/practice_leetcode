package practice_leetcode

var m = map[int]int{0: 0, 1: 1, 2: 2}
var maxN = 2

func climbStairs(n int) int {
	if n > maxN {
		for i := maxN + 1; i <= n; i++ {
			m[i] = m[i-1] + m[i-2]
		}
	}
	return m[n]
}
