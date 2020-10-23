package solution

import "sort"

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	dp := make([][]int, len(nums))
	rsu := []int{nums[0]}
	dp[0] = []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if len(dp[j]) >= len(dp[i]) && nums[i]%nums[j] == 0 {
				dp[i] = make([]int, len(dp[j])+1)
				copy(dp[i], dp[j])
				dp[i][len(dp[i])-1] = nums[i]
			}
		}
		if dp[i] == nil || len(dp[i]) == 0 {
			dp[i] = []int{nums[i]}
			continue
		}
		if len(dp[i]) > len(rsu) {
			rsu = dp[i]
		}
	}
	return rsu
}
