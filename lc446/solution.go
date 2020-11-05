package solution

import "sort"

var current, nums_ []int
var rsu int

func numberOfArithmeticSlices(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	nums_ = nums
	current = []int{}
	rsu = 0
	dfs(0)
	return rsu
}

func dfs(startIdx int) {
	if len(current) >= 3 {
		if current[len(current)-2]-current[len(current)-3] != current[len(current)-1]-current[len(current)-2] {
			return
		}
		rsu++
	}
	for i := startIdx; i < len(nums_); i++ {
		current = append(current, nums_[i])
		dfs(i + 1)
		current = current[:len(current)-1]
	}
}
