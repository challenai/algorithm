package solution

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var prev, current, resu0, temp int
	// not rob the first one
	prev = 0
	current = 0
	for i := 1; i < len(nums); i++ {
		temp = current
		if prev+nums[i] > current {
			current = prev + nums[i]
		}
		prev = temp
	}
	resu0 = current
	prev = 0
	current = nums[0]
	for i := 1; i < len(nums)-1; i++ {
		temp = current
		if prev+nums[i] > current {
			current = prev + nums[i]
		}
		prev = temp
	}
	if current > resu0 {
		return current
	}
	return resu0
}
