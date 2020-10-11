package solution

import (
	"sort"
)

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	resu := nums[1] - nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > resu {
			resu = nums[i] - nums[i-1]
		}
	}
	return resu
}
