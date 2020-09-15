package solution

func removeDuplicates(nums []int) int {
	var dulplicateIdx, i int
	var okDul bool
	i = 0
	hash := map[int]int{}
	dulplicateIdx = len(nums)
	for i < dulplicateIdx {
		_, ok := hash[nums[i]]
		if ok {
			okDul = true
			// the wrong condition is here
			// I finally handle it but I dont know why
			for okDul && i < dulplicateIdx {
				dulplicateIdx--
				_, okDul = hash[nums[dulplicateIdx]]
			}
			nums[i], nums[dulplicateIdx] = nums[dulplicateIdx], nums[i]
		}
		hash[nums[i]] = 0
		i++
	}
	return dulplicateIdx
}
