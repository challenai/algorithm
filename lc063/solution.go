package solution

func uniquePathsWithObstacles(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	temp := 1
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// init state
	for i := 0; i < m; i++ {
		if matrix[i][0] == 1 {
			temp = 0
		}
		dp[i][0] = temp
	}
	temp = 1
	for i := 0; i < n; i++ {
		if matrix[0][i] == 1 {
			temp = 0
		}
		dp[0][i] = temp
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i-1][j] == 0 {
				dp[i][j] += dp[i-1][j]
			}
			if matrix[i][j-1] == 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
