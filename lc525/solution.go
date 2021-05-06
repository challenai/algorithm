package solution

func findMaxLength(nums []int) int {
	rsu := 0
	if len(nums) == 0 {
		return 0
	}
	if nums[0] == 0 {
		nums[0] = -1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = nums[i-1] - 1
			continue
		}
		nums[i] += nums[i-1]
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j += 2 {
			if i == 0 {
				if nums[j] == 0 && j+1 > rsu {
					rsu = j + 1
				}
				continue
			}
			if nums[j]-nums[i-1] == 0 && j-i+1 > rsu {
				rsu = j - i + 1
			}
		}
	}
	return rsu
}
