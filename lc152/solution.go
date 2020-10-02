package solution

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lastMin := 0
	lastMax := 0
	resu := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			lastMax = 0
			lastMin = 0
		} else if nums[i] > 0 {
			if lastMax == 0 {
				lastMax = nums[i]
			} else {
				lastMax *= nums[i]
			}
			if lastMin != 0 {
				lastMin *= nums[i]
			}
		} else {
			if lastMin == 0 {
				lastMax = 0
				lastMin = nums[i]
			} else {
				lastMax = nums[i] * lastMin
			}
			if lastMax != 0 {
				lastMin = lastMax * nums[i]
			}
		}
		if resu < lastMax {
			resu = lastMax
		}
	}

	return resu
}
