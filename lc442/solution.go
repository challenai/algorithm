package solution

func findDuplicates(nums []int) []int {
	rsu := []int{}
	hash := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			rsu = append(rsu, nums[i])
			continue
		}
		hash[nums[i]] = false
	}
	return rsu
}
