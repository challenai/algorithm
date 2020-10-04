package solution

func maxProfit(prices []int, k int) int {
	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, 2)
	}

	// init state
	for i := 0; i <= k; i++ {
		dp[i][0] = 0
		dp[i][1] = -prices[0]
	}

	for i := 0; i < len(prices); i++ {
		for j := 0; j <= k; j++ {
			if j > 0 && dp[j-1][1]+prices[i] > dp[j][0] {
				dp[j][0] = dp[j-1][1] + prices[i]
			}
			if dp[j][0]-prices[i] > dp[j][1] {
				dp[j][1] = dp[j][0] - prices[i]
			}
		}
	}

	return dp[k][0]
}
