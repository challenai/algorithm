package solution

import "sort"

func thirdMax(nums []int) int {
	if len(nums) < 3 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	list := make([]int, 3)
	copy(list, nums[:3])
	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	})
	for i := 3; i < len(nums); i++ {
		if nums[i] == list[0] || nums[i] == list[1] {
			continue
		}
		if nums[i] > list[2] || list[1] == list[2] {
			list[2] = nums[i]
			if list[2] > list[1] {
				list[1], list[2] = list[2], list[1]
				if list[1] > list[0] {
					list[0], list[1] = list[1], list[0]
				}
			}
		}
	}
	if list[0] == list[1] || list[1] == list[2] {
		return list[0]
	}
	return list[2]
}
