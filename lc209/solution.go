package solution

func minSubArrayLen(s int, nums []int) int {
	// sliding window~
	if len(nums) == 0 {
		return 0
	}
	if nums[0] >= s {
		return 1
	}
	var left, right, sum int
	sum = nums[1]
	left = 0
	for i := 1; i < len(nums); i++ {
		if 
	}
}
