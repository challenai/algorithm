package solution

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	limit := 2
	dp := [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}

	// buy, first 1 represent have stock
	// bug here before: if I buy the first, my prices is always -prices[0], no matter how many limit I have left, It just doesnt matter as long as we have more than once
	// the code before are as follows:
	// dp[1][0] = -prices[0]
	for i := 0; i <= limit; i++ {
		dp[1][i] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := limit; j > 0; j-- {
			// the bug happens here,
			// the bug code before as follows:
			// dp[1][j] = max(dp[0][j+1]-prices[i], dp[1][j])
			dp[1][j] = max(dp[0][j-1]-prices[i], dp[1][j])
			dp[0][j] = max(dp[0][j], dp[1][j]+prices[i])
		}
	}

	return dp[0][limit]
}

// var prices_ []int

// func maxProfit(prices []int) int {
// 	prices_ = prices
// 	return loop(len(prices)-1, 0, 2)
// }

// // limit represent buy times limit
// func loop(n int, hold int, limit int) int {
// 	if n == 0 {
// 		if hold == 1 {
// 			return -prices_[0]
// 		}
// 		return 0
// 	}
// 	if limit == 0 {
// 		return 0
// 	}
// 	if hold == 1 {
// 		return max(loop(n-1, 1, limit), loop(n-1, 0, limit-1)-prices_[n])
// 	} else {
// 		return max(loop(n-1, 1, limit)+prices_[n], loop(n-1, 0, limit))
// 	}
// }

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
