package solution

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	temp := 0
	for i := 0; i < len(grid); i++ {
		temp += grid[i][0]
		dp[i][0] = temp
	}
	temp = 0
	for i := 0; i < len(grid[0]); i++ {
		temp += grid[0][i]
		dp[0][i] = temp
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
