package practice_leetcode

func generateMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}
	x, y := 0, 0
	v := 1
	for ; n > 0; n -= 2 {
		if n == 1 {
			m[y][x] = v
		}
		for i := 0; i < n-1; i++ {
			m[y][x] = v
			x++
			v++
		}
		for i := 0; i < n-1; i++ {
			m[y][x] = v
			y++
			v++
		}
		for i := 0; i < n-1; i++ {
			m[y][x] = v
			x--
			v++
		}
		for i := 0; i < n-1; i++ {
			m[y][x] = v
			y--
			v++
		}
		x++
		y++
	}
	return m
}
