package solution

func checkSubarraySum(nums []int, k int) bool {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if i == 0 {
				if nums[j]%k == 0 {
					return true
				}
				continue
			}
			if (nums[j]-nums[i-1])%k == 0 {
				return true
			}
		}
	}

	return false
}
