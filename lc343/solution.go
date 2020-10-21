package solution

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	dp := []int{0, 0, 1}
	for i := 3; i <= n; i++ {
		dp = append(dp, i)
		for j := 2; j <= i>>1; j++ {
			dp[len(dp)-1] = max(dp[len(dp)-1], max(dp[j], j)*max(dp[i-j], i-j))
		}
	}
	// for i := 0; i < len(dp); i++ {
	// 	print(dp[i], " ")
	// }
	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
