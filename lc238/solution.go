package solution

func productExceptSelf(nums []int) []int {
	resu := make([]int, len(nums))
	var current int
	current = 1
	for i := 0; i < len(nums); i++ {
		current *= nums[i]
		resu[i] = current
	}
	current = 1
	for i := len(nums) - 1; i > 0; i-- {
		resu[i] = resu[i-1] * current
		current *= nums[i]
	}
	resu[0] = current
	return resu
}
