package solution

func findMedianSortedArrays(nums1, nums2 []int) int {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	var current, last int
	nums := [2][]int{nums1, nums2}
	idxs := [2]int{0, 0}
	for true {
		last = current
		if idxs[0] < len(nums[0]) && idxs[1] < len(nums[1]) {
			// selectMin
			if nums[0][idxs[0]] < nums[1][idxs[1]] {
				idxs[0]++
				current = nums[0][idxs[0]]
			} else {
				idxs[1]++
				current = nums[1][idxs[1]]
			}
		} else if idxs[0] < len(nums[0]) {
			// select(0)
			idxs[0]++
			current = nums[0][idxs[0]]
		} else {
			// select(1)
			idxs[1]++
			current = nums[1][idxs[1]]
		}
		if idxs[0]+idxs[1] == (len(nums[0])+len(nums[1]))/2 {
			if (len(nums[0])+len(nums[1]))%2 == 1 {
				return current
			} else {
				return (current + last) / 2
			}
		}
	}

	return 0
}
