package solution

func countSmaller(nums []int) []int {
	rsu := make([]int, len(nums))
	st := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(st) > 0 && nums[i] < st[len(st)-1] {
			st = st[:len(st)-1]
		}
		rsu[i] = len(st)
		st = append(st, nums[i])
	}
	return rsu
}
