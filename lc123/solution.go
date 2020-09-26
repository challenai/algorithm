package solution

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}

	// buy, first 1 represent have stock
	dp[1][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 3; j++ {
			if j <= 1 {
				dp[1][j] = max(dp[0][j+1]-prices[i], dp[1][j])
			}
			dp[0][j] = max(dp[0][j], dp[1][j]+prices[i])
		}
	}
	println(dp[0][2])

	return dp[0][2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	if b > a {
		a = b
	}
	if c > a {
		a = c
	}
	return a
}
