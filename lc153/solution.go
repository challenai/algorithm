package solution

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}
	var low, high, mid int
	low = 0
	high = len(nums) - 1
	for low < high-1 {
		mid = (low + high) / 2
		if nums[low] > nums[high] {
			low = mid
		} else {
			high = mid
		}
	}
	return nums[low]
}
