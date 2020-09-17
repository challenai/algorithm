package solution

func firstMissingPositive(nums []int) int {
	hasLen := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == len(nums) {
			hasLen = true
			continue
		}
		for nums[i] > 0 && nums[i] < len(nums) && nums[nums[i]] != nums[i] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}

	}

	for i := 1; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	if hasLen {
		return len(nums) + 1
	}
	return len(nums)
}
