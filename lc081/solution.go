package solution

func search(nums []int, target int) bool {

	var left, right, mid int
	left = 0
	right = len(nums) - 1
	for left < right-1 {
		mid = (left + right) / 2
		if nums[mid] > nums[left] {
			left = mid
		} else {
			right = mid
		}
	}
	if nums[left] == nums[right] {
		return nums[left] == target
	}
	if target >= nums[0] {
		right = left
		left = 0
	} else {
		left = right
		right = len(nums) - 1
	}

	for left < right-1 {
		mid = (left + right) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			left = mid
		} else {
			right = mid
		}
	}
	return nums[left] == target || nums[right] == target
}
