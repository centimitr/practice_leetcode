package practice_leetcode

func uniquePaths(m int, n int) int {
	rows := make([][]int, m+1)
	for i := range rows {
		rows[i] = make([]int, n+1)
	}
	rows[0][1] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			rows[i][j] = rows[i-1][j] + rows[i][j-1]
		}
	}
	return rows[m][n]
}
