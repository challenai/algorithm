package solution

func lastStoneWeight(nums []int) int {
	hp := make([]int, len(nums))
	right := -1
	var ptr int
	for i := 0; i < len(nums); i++ {
		right++
		hp[right] = nums[i]
		ptr = right
		for ptr > 0 && hp[ptr] > hp[(ptr-1)>>1] {
			hp[ptr], hp[(ptr-1)>>1] = hp[(ptr-1)>>1], hp[ptr]
			ptr = (ptr - 1) >> 1
		}
	}
	current := [2]int{0, 0}
	for right > 0 {
		// for i := 0; i <= right; i++ {
		// 	print(hp[i], " ")
		// }
		// println()
		current[0] = hp[0]
		reheapify(hp, right)
		right--
		current[1] = hp[0]
		reheapify(hp, right)
		// right--
		hp[right] = current[0] - current[1]
		// println(current[0], current[1], hp[right])
		ptr = right
		for ptr > 0 && hp[ptr] > hp[(ptr-1)>>1] {
			hp[ptr], hp[(ptr-1)>>1] = hp[(ptr-1)>>1], hp[ptr]
			ptr = (ptr - 1) >> 1
		}
	}

	return hp[0]
}

func reheapify(hp []int, right int) {
	ptr := 0
	hp[ptr] = hp[right]
	for (ptr<<1)+2 < right {
		if hp[ptr<<1+1] < hp[ptr<<1+2] {
			hp[ptr], hp[ptr<<1+2] = hp[ptr<<1+2], hp[ptr]
			ptr = ptr<<1 + 2
		} else {
			hp[ptr], hp[ptr<<1+1] = hp[ptr<<1+1], hp[ptr]
			ptr = ptr<<1 + 1
		}
	}
	if ptr<<1+2 == right {
		if hp[ptr] < hp[ptr<<1+1] {
			hp[ptr], hp[ptr<<1+1] = hp[ptr<<1+1], hp[ptr]
		}
	}
}
