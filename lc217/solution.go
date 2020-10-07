package solution

func containsDuplicate(nums []int) bool {
	// I know we have bitwise O(1) sc solution, but I dont know how...
	// so I just use a simple hash
	hash := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			return true
		}
		hash[nums[i]] = false
	}
	return false
}
