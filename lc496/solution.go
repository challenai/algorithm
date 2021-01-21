package solution

func NextGreaterElement(nums1, nums2 []int) []int {
	hash := map[int]int{}
	rsu := make([]int, len(nums1))
	st := []int{}
	for i := 0; i < len(nums2); i++ {
		for len(st) > 0 && nums2[i] > nums2[st[len(st)-1]] {
			// println(nums2[st[len(st)-1]], " ", nums2[i])
			if len(st) > 0 {
				hash[nums2[st[len(st)-1]]] = nums2[i]
			}
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	for i := 0; i < len(nums1); i++ {
		if v, ok := hash[nums1[i]]; ok {
			rsu[i] = v
			continue
		}
		rsu[i] = -1
	}
	return rsu
}
