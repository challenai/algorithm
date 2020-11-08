package solution

func totalHammingDistance(nums []int) int {
	rsu := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			temp := nums[i] ^ nums[j]
			for temp > 0 {
				if temp&1 == 1 {
					rsu++
				}
				temp >>= 1
			}
		}
	}
	return rsu
}
