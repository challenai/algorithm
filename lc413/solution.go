package solution

func numberOfArithmeticSlices(nums []int) int {
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = i
	}
	rsu := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			i_ := i
			j_ := j
			idx, ok := hash[2*nums[j_]-nums[i_]]
			for ok {
				i_ = j_
				j_ = idx
				rsu++
				idx, ok = hash[2*nums[j_]-nums[i_]]
			}
		}
	}

	return rsu
}
