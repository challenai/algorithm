package solution

func isSelfCrossing(x []int) bool {
	// basic idea look like this,
	// check out if current pos is located in a specific range
	ranges := [4]int{0, 0, 0, 0}
	pos := [2]int{0, 0}
	for i := 0; i < len(x); i++ {
		if i%4 == 0 {
			pos[1] += x[i]
		} else if i%4 == 1 {
			pos[0] -= x[i]
		} else if i%4 == 2 {
			pos[1] -= x[i]
		} else {
			pos[0] += x[i]
		}

		if i%4 == 0 || i%4 == 2 {
			if i != 0 && pos[0] <= ranges[1] && pos[0] >= ranges[0] {
				return true
			}
			if pos[0] > ranges[1] {
				ranges[1] = pos[0]
			} else if pos[0] < ranges[0] {
				ranges[0] = pos[0]
			}
		} else {
			if i != 1 && pos[1] <= ranges[3] && pos[1] >= ranges[2] {
				return true
			}
			if pos[1] > ranges[3] {
				ranges[3] = pos[1]
			} else if pos[1] < ranges[2] {
				ranges[2] = pos[1]
			}
		}
	}
	return false
}
