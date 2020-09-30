package solution

func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for j := 3; j <= n; j++ {
		for i := 1; i < j-1; i++ {
			dp[j] = dp[i] * dp[j-i-1]
		}
		dp[j] += 2 * dp[j-1]
	}
	return dp[n]
}
