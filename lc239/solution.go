package solution

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	if len(nums) == 0 {
		return nums
	}
	// idea: use a queue to store the max data
	wndw := []int{nums[0]}
	resu := []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i] < wndw[0] {
			wndw = append(wndw, nums[i])
			if len(wndw) > k {
				wndw = wndw[1:]
			}
		} else {
			wndw = []int{nums[i]}
		}
		if i >= k-1 {
			resu = append(resu, wndw[0])
		}
	}
	return resu
}
