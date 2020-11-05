package solution

func findDisappearedNumbers(nums []int) []int {
	rsu := []int{}
	maxN := len(nums)
	var temp int
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = maxN + 1
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			temp = -nums[i]
		} else {
			temp = nums[i]
		}
		if temp >= 1 && temp <= maxN {
			if nums[temp-1] > 0 {
				nums[temp-1] = -nums[temp-1]
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			rsu = append(rsu, i+1)
		}
	}

	return rsu
}
