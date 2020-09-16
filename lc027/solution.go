package solution

func removeElement(nums []int, val int) int {
  removeIdx := len(nums)
  for i := len(nums) -1; i >= 0; i-- {
    if nums[i] == val {
      removeIdx--
      nums[i], nums[removeIdx] = nums[removeIdx], nums[i]
    }
  }

  return removeIdx
}
