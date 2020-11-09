package solution

import "sort"

func heightChecker(nums []int) int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	rsu := 0
	for i := 0; i < len(sorted); i++ {
		if sorted[i] != nums[i] {
			rsu++
		}
	}
	return rsu
}
