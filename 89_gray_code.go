package practice_leetcode

func grayCode(n int) []int {
	code := make([]int, 1<<uint(n))
	for i := 0; i < 1<<uint(n); i++ {
		code[i] = i ^ (i / 2)
	}
	return code
}
