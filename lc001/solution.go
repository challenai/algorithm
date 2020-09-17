package solution

func twoSum(nums []int, target int) []int {
	// inspire:
	// use a hash to get a O(1) complexity to search. use minus instead of sum to do == op
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if idx, ok := hash[target-nums[i]]; ok {
			return []int{idx, i}
		}
		hash[nums[i]] = i
	}
	return []int{}
}
