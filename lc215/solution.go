package solution

func findKthLargest(nums []int, k int) int {
	if k < 0 || k > len(nums) {
		return 0
	}
	var resuIdx, resu int
	for k > 0 {
		resuIdx = 0
		for i := 1; i < len(nums); i++ {
			if nums[i] >= nums[resuIdx] {
				resuIdx = i
			}
		}
		resu = nums[resuIdx]
		if resuIdx != len(nums) {
			copy(nums[resuIdx:], nums[resuIdx+1:])
		}
		nums = nums[:len(nums)-1]
		k--
	}
	return resu
}
