package solution

import "sort"

func hIndex(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] < i+1 {
			break
		}
	}
	if nums[len(nums)-1] >= i+1 {
		i++
	}
	return i
}
