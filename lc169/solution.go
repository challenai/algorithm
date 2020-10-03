package solution

func majorityElement(nums []int) int {
	// I have a really good idea. I destory every different pair
	// the last single majority word is result
	// because always use hash if so boring, just for fun :)
	if len(nums) == 0 {
		return 0
	}
	quorum := [2]int{0, 1}
	for i := 1; i < len(nums); i++ {
		if quorum[0] == -1 {
			quorum[0] = i
			quorum[1] = 1
		} else if quorum[0] != nums[i] {
			if quorum[1] > 1 {
				quorum[1]--
			} else {
				quorum[0] = -1
			}
		} else {
			quorum[1]++
		}
	}

	if quorum[0] == -1 {
		return 0
	}

	return nums[quorum[0]]
}
