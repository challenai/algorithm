package solution

func find132pattern(nums []int) bool {
	st := []int{}
	for i := 0; i < len(nums); i++ {
		if len(st) > 0 && nums[i] < st[len(st)-1] {
			st = st[:len(st)-1]
			if len(st) > 0 {
				return true
			}
		}
		st = append(st, nums[i])
	}
	return false
}
