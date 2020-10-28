package solution

import "sort"

func reconstructQueue(nums [][]int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i][0] > nums[j][0] {
			return true
		}
		if nums[i][0] < nums[j][0] {
			return false
		}
		if nums[i][1] < nums[j][1] {
			return true
		}
		return false
	})
	if nums[0][1] > 0 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		if nums[i][1] < i {
			// println(nums[i][1])
			// temp := nums[i]
			// if nums[i][1] > 0 {
			// 	copy(nums[nums[i][1]-1:i], nums[nums[i][1]:i+1])
			// 	nums[nums[i][1]-1] = temp
			// } else {
			// 	copy(nums[:i], nums[:i+1])
			// 	nums[0] = temp
			// }
			temp := nums[i][1]
			for j := i; j > temp; j-- {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}

	// for i := 0; i < len(nums); i++ {
	// 	print(nums[i][0], " ", nums[i][1], " | ")
	// }
	// println()

	return nums
}
