package solution

func findDuplicate(nums []int) int {
	// ok I dont know the right bitwise solution
	// but Why I should use bitwise handle such kind of invalueable problem...
	// hash O(n) tc solution are as follows:
	hash := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			return nums[i]
		} else {
			hash[nums[i]] = false
		}
	}
	return -1
}
