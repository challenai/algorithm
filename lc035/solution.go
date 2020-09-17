package solution

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if nums[0] >= target {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	var ptr, left, right int
	left = 0
	right = len(nums) - 1
	for left < right-1 {
		ptr = (left + right) / 2
		if nums[ptr] == target {
			return ptr
		} else if nums[ptr] > target {
			right = ptr
		} else {
			left = ptr
		}
	}
	return right
}
