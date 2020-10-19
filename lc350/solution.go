package solution

func intersect(nums0, nums1 []int) []int {
	resu := []int{}
	hash := map[int]int{}
	for i := 0; i < len(nums0); i++ {
		if _, ok := hash[nums0[i]]; ok {
			hash[nums0[i]]++
		} else {
			hash[nums0[i]] = 1
		}
	}
	for i := 0; i < len(nums1); i++ {
		if _, ok := hash[nums1[i]]; ok {
			if hash[nums1[i]] == 1 {
				delete(hash, nums1[i])
			} else {
				hash[nums1[i]]--
			}
			resu = append(resu, nums1[i])
		}
	}
	return resu
}
