package solution

func sortColors(nums []int) {
	var left, right int
	left = -1
	right = len(nums)
	for i := 0; i < right; i++ {
		if nums[i] == 0 {
			left++
			nums[left], nums[i] = nums[i], nums[left]
		} else if nums[i] == 2 {
			right--
			nums[right], nums[i] = nums[i], nums[right]
		}
	}
	// for i := 0; i < len(nums); i++ {
	// print(nums[i], " ")
	// }
}
