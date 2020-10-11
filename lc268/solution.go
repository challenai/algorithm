package solution

func missingNumber(nums []int) int {
	hasLast := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == len(nums) {
			hasLast = true
		} else if nums[i] != i {
			for nums[i] != i {
				if nums[i] == len(nums) {
					break
				}
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			}
		}
	}
	if !hasLast {
		return len(nums)
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return -1
}
