package solution

func findMaxLength(nums []int) int {
	rsu := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 1 {
			nums[i] = -1
		}
	}
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j += 2 {
			if i == 0 {
				if nums[j] == 0 {
					rsu = max(rsu, j+1)
				}
				continue
			}
			if nums[j]-nums[i-1] == 0 {
				rsu = max(rsu, j-i+1)
			}
		}
	}
	return rsu
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
