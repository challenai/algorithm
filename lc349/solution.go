package solution

func intersection(nums1, nums2 []int) []int {
	hash := map[int]bool{}
	resu := []int{}
	for i := 0; i < len(nums1); i++ {
		hash[nums1[i]] = false
	}
	for i := 0; i < len(nums2); i++ {
		if _, ok := hash[nums2[i]]; ok {
			resu = append(resu, nums2[i])
			delete(hash, nums2[i])
		}
	}
	return resu
}
