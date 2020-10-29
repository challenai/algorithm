package solution

func addStrings(nums1, nums2 string) string {
	resu := ""
	var needPlusOne, temp byte
	needPlusOne = 0
	for i := 0; i < max(len(nums1), len(nums2)); i++ {
		if i < len(nums1) && i < len(nums2) {
			temp = nums1[len(nums1)-1-i] + nums2[len(nums2)-1-i] + needPlusOne - 48
		} else if i < len(nums1) {
			temp = nums1[len(nums1)-1-i] + needPlusOne
		} else {
			temp = nums2[len(nums2)-1-i] + needPlusOne
		}
		needPlusOne = 0
		if temp > 57 {
			temp -= 10
			needPlusOne = 1
		}
		resu = string(temp) + resu
	}
	// 0 = 48
	// 9 = 57
	return resu
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
