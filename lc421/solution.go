package solution

func findMaximumXOR(nums []int) int {
	resu := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]^nums[j] > resu {
				resu = nums[i] ^ nums[j]
			}
		}
	}
	return resu
}
