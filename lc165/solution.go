package solution

func compareVersion(v1, v2 string) int {
	ptrs := []int{0, 0}
	for ptrs[0] < len(v1) && ptrs[1] < len(v2) {
		if temp := compareNextSeg(v1, v2, ptrs); temp != 0 {
			return temp
		}
	}
	if ptrs[0] == len(v1) && ptrs[1] == len(v2) {
		return 0
	}
	if ptrs[0] != len(v1) && searchEffectiveNum(v1, ptrs[0]) {
		return 1
	}
	if ptrs[1] != len(v2) && searchEffectiveNum(v2, ptrs[1]) {
		return -1
	}
	return 0
}

func compareNextSeg(v1, v2 string, ptrs []int) int {
	for ptrs[0] < len(v1) && v1[ptrs[0]] == '0' {
		ptrs[0]++
	}
	for ptrs[1] < len(v2) && v2[ptrs[1]] == '0' {
		ptrs[1]++
	}
	for ptrs[0] < len(v1) && ptrs[1] < len(v2) {
		if v1[ptrs[0]] == '.' || v2[ptrs[1]] == '.' {
			if v1[ptrs[0]] != '.' {
				return 1
			}
			if v2[ptrs[1]] != '.' {
				return -1
			}
			ptrs[0]++
			ptrs[1]++
			return 0
		}
		if v1[ptrs[0]] != v2[ptrs[1]] {
			if v1[ptrs[0]] > v2[ptrs[1]] {
				return 1
			} else {
				return -1
			}
		}
		ptrs[0]++
		ptrs[1]++
	}
	return 0
}

func searchEffectiveNum(v string, ptr int) bool {
	for ptr < len(v) {
		if v[ptr] != '.' && v[ptr] != '0' {
			return true
		}
		ptr++
	}
	return false
}
