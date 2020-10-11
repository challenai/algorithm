package solution

func singleNumber(nums []int) []int {
	hash := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			delete(hash, nums[i])
		} else {
			hash[nums[i]] = false
		}
	}
	resu := []int{}
	for k := range hash {
		resu = append(resu, k)
	}
	return resu
}
