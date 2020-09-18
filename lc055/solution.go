package solution

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		if !dp[i] {
			return false
		}
		for j := 1; j <= nums[i]; j++ {
			if j+i < len(dp) {
				dp[i+j] = true
			}
		}
	}
	return true
}
