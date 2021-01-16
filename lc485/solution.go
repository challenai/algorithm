package solution

func FindMaxConsecutiveOnes(nums []int) int {
	rsu := 0
	current := 0
	if len(nums) == 0 {
		return rsu
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			current = 0
			continue
		}
		current++
		if current > rsu {
			rsu = current
		}
	}
	return rsu
}
