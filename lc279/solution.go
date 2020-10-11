package solution

func numSquares(n int) int {
	if n < 1 {
		return 0
	}
	options := []int{}
	for i := 1; i*i <= n; i++ {
		if i*i <= n {
			if i*i == n {
				return 1
			}
			options = append(options, i*i)
		}
	}
	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = i
	}
	for i := 0; i < len(options); i++ {
		for j := 2; j <= n; j++ {
			if options[i] == j {
				dp[j] = 1
				continue
			}
			if j-options[i] >= 0 && dp[j] > dp[j-options[i]]+1 {
				dp[j] = dp[j-options[i]] + 1
			}
		}
	}

	return dp[n]
}
