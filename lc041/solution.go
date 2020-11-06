package solution

func firstMissingPositive2(nums []int) int {
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

func firstMissingPositive(nums []int) int {
	// a really great idea is use mark, 0, 1, -1 are all great choice for mark
	var temp int
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 2
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			temp = -nums[i]
		} else {
			temp = nums[i]
		}
		if temp <= len(nums) {
			nums[temp-1] = -nums[temp-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
