package solution

func containsNearbyDuplicate(nums []int, k int) bool {
	// idea: mantain a len(k) sliding window
	wndw := map[int]bool{}
	var left int
	left = 0
	for i := 0; i < len(nums); i++ {
		if i-left > k {
			delete(wndw, nums[left])
			left++
		}
		if _, ok := wndw[nums[i]]; ok {
			return true
		} else {
			wndw[nums[i]] = false
		}
	}
	return false
}
