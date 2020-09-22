package solution

func largestRectangleArea(nums []int) int {
	// idea: its a really difficult problem,
	// however, ocasionally I learn monostack these days
	// I think monostack is really suitable for the kind of problem which wanna you find the most close larger or smaller ones in both sides
	// A practical trick could be you push the index into stack instead of specific height.
	// at last it's a ascent st.
	var resu, temp int
	resu = 0
	st := []int{-1}
	for i := 0; i < len(nums); i++ {
		if len(st) > 1 && nums[i] < nums[st[len(st)-1]] {
			// pop el
			for len(st) > 1 && nums[i] < nums[st[len(st)-1]] {
				temp = st[len(st)-1]
				st = st[:len(st)-1]
				resu = max(resu, (i-st[len(st)-1]-1)*nums[temp])
			}
		}
		// push el, just add a push in for loop,
		// bug here before, I put push action in a else branch
		// but EVERY element should be pushed to st once!!!
		st = append(st, i)
	}

	for len(st) > 1 {
		// pop el
		temp = st[len(st)-1]
		st = st[:len(st)-1]
		resu = max(resu, (temp-st[len(st)-1])*nums[temp])
	}

	return resu
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
