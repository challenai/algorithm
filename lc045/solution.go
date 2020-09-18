package solution

func jump(nums []int) int {
	dp := make([]int, len(nums))

	// init state
	for i := 1; i < len(nums); i++ {
		dp[i] = len(nums)
	}
	dp[0] = 0

	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j < len(nums) {
				dp[i+j] = min(dp[i+j], nums[i]+1)
			}
		}
	}

	return dp[len(nums)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
