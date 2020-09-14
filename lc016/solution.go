package solution

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	// idea: really similar to the foreign problme...
	var resu, left, right, current int
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	resu = nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		left = i + 1
		right = len(nums) - 1
		for left < right {
			current = nums[i] + nums[left] + nums[right]
			if current-target == 0 {
				return target
			}
			if current > target {
				right--
			} else {
				left++
			}
			if abs(current-target) < abs(resu-target) {
				resu = current
			}
		}
	}
	return resu
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
