package solution

func singleNumber(nums []int) int {
	resu := 0
	for i := 0; i < len(nums); i++ {
		resu ^= nums[i]
	}
	return resu
}
