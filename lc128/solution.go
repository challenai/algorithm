package solution

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	resu := 1
	hash := map[int][2]int{}
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = [2]int{nums[i], nums[i]}
	}

	for _, ranges := range hash {
		_, ok := hash[ranges[0]-1]
		for ok {
			ranges[0]--
			delete(hash, ranges[0])
			_, ok = hash[ranges[0]-1]
		}
		_, ok = hash[ranges[1]+1]
		for ok {
			ranges[1]++
			delete(hash, ranges[1])
			_, ok = hash[ranges[1]+1]
		}
		if ranges[1]-ranges[0]+1 > resu {
			resu = ranges[1] - ranges[0] + 1
		}
	}

	return resu
}
