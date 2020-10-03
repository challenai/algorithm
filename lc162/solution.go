package solution

func findPeakElement(nums []int) int {
	var low, high, mid int
	low = -1
	high = len(nums) - 1
	for low < high {
		mid = low + (high-low)/2
		// 如果中间值比后面大
		// 说明在下坡， 峰值就在左边
		// 因为下坡可能会一直下
		// 从左边找个上坡，就好了
		// 最后停下来的时候高低位必然相差1
		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
