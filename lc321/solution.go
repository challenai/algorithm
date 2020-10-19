package solution

func maxNumber(nums1, nums2 []int, k int) []int {
	resu := []int{}
	var ptr1, ptr2 int
	ptr1 = 0
	ptr2 = 0
	for k > 0 {
		current := 0
		ptr1_ := ptr1
		ptr2_ := ptr2
		maxLen1 := len(nums1)
		if len(nums2)-ptr2 < k {
			k = maxLen1 - (len(nums2) - ptr2)
		}
		for i := ptr1; i < maxLen1; i++ {
			if nums1[i] > current {
				current = nums1[i]
				ptr1_ = i
			}
		}
		maxLen2 := len(nums2)
		if len(nums1)-ptr1 < k {
			k = maxLen2 - (len(nums1) - ptr1)
		}
		for i := 0; i < maxLen2; i++ {
			if nums2[i] > current {
				current = nums2[i]
				ptr2_ = i
			}
		}
		resu = append(resu, current)
		if ptr1_ >= len(nums1) {
			ptr2 = ptr2_
			ptr2++
		} else if ptr2_ >= len(nums2) {
			ptr1 = ptr1_
			ptr1++
		} else {
			if nums1[ptr1_] > nums2[ptr2_] {
				ptr1 = ptr1_
				ptr1++
			} else {
				ptr2 = ptr2_
				ptr2++
			}
		}
		k--
	}
	for i := 0; i < len(resu); i++ {
		print(resu[i], " ")
	}
	print()
	return resu
}
