package solution

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	var low1, low2 int
	low1 = 0
	low2 = -1
	for i := 1; i < len(nums); i++ {
		if low2 != -1 && nums[i] > nums[low2] {
			return true
		}
		if nums[i] > nums[low1] {
			if low2 == -1 || nums[low2] > nums[i] {
				low2 = i
			}
		}
		if nums[i] < nums[low1] {
			low1 = i
		}
	}
	return false
}
