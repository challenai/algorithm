package solution

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := [2]int{1, 1}
	current := [2]int{0, 0}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[current[1]] {
			current[1] = i
			dp[1]++
			dp[0], dp[1] = dp[1], dp[0]
			current[0], current[1] = current[1], current[0]
		} else if nums[i] < nums[current[0]] {
			current[0] = i
			dp[0]++
			dp[0], dp[1] = dp[1], dp[0]
			current[0], current[1] = current[1], current[0]
		}
	}
	if dp[0] > dp[1] {
		return dp[0]
	}
	return dp[1]
}
