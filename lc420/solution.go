package solution

func strongPasswordChecker(s string) int {
	// insert_cnt, insert_resuable
	// delete_cnt, delete_resuable
	// replace_cnt, replace_resuable
	state := [6]int{0, 0, 0, 0, 0, 0}
	// condition 1
	if len(s) < 6 {
		state[0] += 6 - len(s)
		state[1] = state[0]
	}
	if len(s) > 20 {
		state[2] -= len(s) - 20
		state[3] = state[2]
	}

	// condition2
	condition2 := [3]int{0, 0, 0}
	for i := 0; i < len(s); i++ {
		if s[i] > 'a' && s[i] < 'z' {
			condition2[0] = 1
			continue
		}
		if s[i] > 'A' && s[i] < 'Z' {
			condition2[1] = 1
			continue
		}
		if s[i] > '0' && s[i] < '9' {
			condition2[2] = 1
		}
	}
	condition2Charge := condition2[0] + condition2[1] + condition2[2]
	if state[1] > 0 {
		condition2Charge -= state[1]
		if condition2Charge >= 0 {
			// state[1] = 0
			state[0] += condition2Charge
		} else {
			// state[1] = -condition2Charge
		}
	}

	return rsu
}

func setState() {}
