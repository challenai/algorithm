package solution

func minSubArrayLen(s int, nums []int) int {
	// sliding window~
	if len(nums) == 0 {
		return 0
	}
	if nums[0] >= s {
		return 1
	}
	resu := 0
	sum := nums[1]
	left := 0
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		for sum >= s {
			if resu == 0 || resu > (i-left+1) {
				resu = i - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	return resu
}
