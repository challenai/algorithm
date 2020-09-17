package solution

func maxArea(nums []int) int {
	// idea: brute force is really easy to implement, but it cost a O(n^2) time.
	// mono stack is really difficult to apply in real problem, It takes me about 10 min to make clean what a mono stack has done to promise we can get the smaller one in both sides.
	// its something like this, you keep the whole stack in a increase of decrease order.
	// to ensure this, everytime you push things to stack, use need to drain the larger ones(asume that you use a decrease stack)
	var resu int
	resu = 0
	// decrease st
	st := []int{}
	for i := 0; i < len(nums); i++ {
		for len(st) > 0 && nums[i] > nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			resu = max(resu, (nums[i]-nums[st[len(st)-1]])*(i-st[len(st)-1]))
		}
		st = append(st, i)
	}
	for len(st) > 0 {
		st = st[:len(st)-1]
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
