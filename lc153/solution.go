package solution

func findMin(nums []int) int {
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
