package solution

func nextGreaterElements(nums []int) []int {
	sz := len(nums)
	nums_ := make([]int, sz<<1)
	copy(nums_[:sz], nums)
	copy(nums_[sz:], nums)
	st := []int{}
	rsu := make([]int, sz)
	for i := 0; i < sz; i++ {
		rsu[i] = -1
	}
	for i := 0; i < sz<<1; i++ {
		for len(st) != 0 && nums_[i] > nums[st[len(st)-1]] {
			println(i, st[len(st)-1])
			rsu[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		st = append(st, i%sz)
	}
	for i := 0; i < len(rsu); i++ {
		print(rsu[i], " ")
	}

	return rsu
}
