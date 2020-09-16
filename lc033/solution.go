package solution

func search(nums []int, target int) int {
	// idea: find the rotate pivot then search the number
	var left, right, mid int
	left = 0
	right = len(nums) - 1
	for left < right-1 {
		mid = (left + right) / 2
		if nums[mid] < nums[left] {
			right = mid
		} else {
			left = mid
		}
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
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid
		} else {
			right = mid
		}
	}
	println(left)
	println(right)
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}
