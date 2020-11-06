package solution

func minMoves(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	rsu := 0
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		rsu += maxNum - nums[i]
	}
	return rsu
}
