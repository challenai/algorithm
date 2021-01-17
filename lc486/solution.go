package solution

func PredictTheWinner(nums []int) bool {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}
	for i := 0; i < len(nums); i++ {
		dp[i][i] = nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][len(nums)-1] > 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
