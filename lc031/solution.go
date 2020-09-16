package solution

import "sort"

func nextPermutation(nums []int) {
  reverseIdx := -1
  for i := 0; i < len(nums)-1; i++ {
    if nums[i] < nums[i+1] {
       reverseIdx = i
    }
  }
  if reverseIdx == -1 {
    sort.Slice(nums, func(a, b int) bool {
      return a > b
    })
    return
  }
  nums[reverseIdx], nums[reverseIdx+1] = nums[reverseIdx+1], nums[reverseIdx]
}
