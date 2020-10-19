package solution

func countRangeSum(nums []int, lower, upper int) int {
	// I cant catch up a O(n) solution
	// I can just find all possible combination with a sum array
	if len(nums) == 0 {
		return 0
	}
	rsu := 0
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			temp := nums[j]
			if i != 0 {
				temp -= nums[i-1]
			}
			if temp <= upper && temp >= lower {
				rsu++
			}
		}
	}
	return rsu
}
