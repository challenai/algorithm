package solution

func moveZeroes(nums []int) {
	endPtr := len(nums)
	for i := 0; i < endPtr; i++ {
		if nums[i] == 0 {
			endPtr--
			copy(nums[i:], nums[i+1:])
			nums[endPtr] = 0
		}
	}
}
