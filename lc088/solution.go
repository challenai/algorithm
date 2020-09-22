package solution

func merge(nums1 []int, m int, nums2 []int, n int) {
	var idx1, idx2 int
	for idx1 < m || idx2 < n {
		if (idx1 < m && idx2 < n && nums1[idx1] <= nums2[idx2]) || (idx2 >= n) {
			idx1++
		} else {
			insert(nums1, idx1, nums2[idx2])
			idx2++
			idx1++
			m++
		}
	}
	// for i := 0; i < len(nums1); i++ {
	// 	print(nums1[i], " ")
	// }
}

func insert(nums []int, idx, num int) {
	for i := len(nums) - 1; i > idx; i-- {
		nums[i] = nums[i-1]
	}
	nums[idx] = num
}
