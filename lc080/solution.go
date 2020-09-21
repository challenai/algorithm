package solution

func removeDulplicate(nums []int) int {
	var counter, right int
	counter = 1
	right = len(nums)
	for i := 1; i < right; i++ {
		if nums[i] == nums[i-1] {
			if counter == 1 {
				counter++
			} else {
				for nums[i] == nums[i-1] && i < right {
					print(nums[i], " ..")
					right--
					bubbleToRight(nums, i, right)
				}
				counter = 1
			}
		} else {
			counter = 1
		}
	}

	return right
}

func bubbleToRight(nums []int, idx, right int) {
	for idx < right {
		nums[idx], nums[idx+1] = nums[idx+1], nums[idx]
		idx++
	}
}
