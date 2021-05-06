package solution

func getMinDistance(nums []int, target int, start int) int {
  // [1,2,<3>,4,5]
  var left, right int
  left = start
  right = start
  for left >= 0 || right < len(nums) {
    if left >= 0 {
      if nums[left] == target {
        return start - left
      }
      left--
    }
    if right < len(nums) {
      if nums[right] == target {
        return right - start
      }
      right++
    }
  }
  return -1
}
