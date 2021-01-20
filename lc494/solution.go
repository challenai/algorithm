package solution

func FindTargetSumWays(nums []int, target int) int {
	var loop func(int)
	current := 0
	rsu := 0
	loop = func(i int) {
		if i == len(nums) {
			if current == target {
				rsu++
			}
			return
		}
		current += nums[i]
		loop(i + 1)
		current -= nums[i]
		current += -nums[i]
		loop(i + 1)
		current -= -nums[i]
	}
	loop(0)
	return rsu
}
