package solution

func twoSum(nums []int, target int) []int {
	resu := []int{0, len(nums) - 1}
	for resu[0] < resu[1] {
		if nums[resu[0]]+nums[resu[1]] == target {
			resu[0]++
			resu[1]++
			return resu
		} else if nums[resu[0]]+nums[resu[1]] > target {
			resu[1]--
		} else {
			resu[0]++
		}
	}
	resu[0] = -1
	resu[1] = -1
	return resu
}
