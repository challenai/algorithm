package solution

func searchRange(nums []int, target int) []int {
	var left, right, ptr int
	left = 0
	right = len(nums) - 1
	for left < right-1 {
		ptr = (left + right) / 2
		if nums[ptr] == target {
			break
		} else if nums[ptr] > target {
			right = ptr
		} else {
			left = ptr
		}
	}
	if nums[ptr] != target {
		return []int{-1, -1}
	}
	left = ptr
	right = ptr
	for left > 0 && nums[left-1] == target {
		left--
	}
	for right < len(nums)-1 && nums[right+1] == target {
		right++
	}
	return []int{left, right}
}
