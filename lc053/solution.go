package solution

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var resu, current int
	resu = nums[0]
	current = nums[0]
	for i := 1; i < len(nums); i++ {
		current = max(current+nums[i], nums[i])
		resu = max(current, resu)
	}

	return resu
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
