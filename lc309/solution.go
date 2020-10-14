package solution

func maxProfit(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := [2][2]int{}
	dp[0][0] = 0
	dp[1][0] = -nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[0][0]-nums[i] > dp[1][0] {
			dp[1][0] = dp[0][0] - nums[i]
		}
		// dp[0][0] = dp[0][0]
		dp[0][1] = dp[1][0] + nums[i]
	}
	return max(dp[0][0], dp[0][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
