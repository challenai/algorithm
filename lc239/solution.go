package solution

func maxSlidingWindow(nums []int, k int) []int {
  // keep window asending or desending so that we just need to compare the first element to find maximum or minimum of windows.
  rsu := []int{}
  if len(nums) == 0 {
    return rsu
  }
  if k <= 1 {
    return nums
  }
  window := []int{0}
  for i := 1; i < len(nums); i++ {
    // remove overflow elements at beginning
    if window[0] < i - k + 1{
      window = window[1:]
    }
    // remove smaller elements in list so that our window is purely descending.
    for len(window) > 0 && nums[i] >= nums[window[len(window)-1]] {
      window = window[:len(window)-1]
    }
    window = append(window, i)
    // if we our index arrived at where we need to push result (at least k elements)
    if i >= k-1 {
      rsu = append(rsu, nums[window[0]])
    }
  }
  return rsu
}
