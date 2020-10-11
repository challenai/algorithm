package solution

func firstBadVersion(n, bad int) int {
	var low, high, mid int
	low = 0
	high = n
	for low < high-1 {
		mid = low + (high-low)>>1
		if isBadVersion(mid, bad) {
			high = mid
		} else {
			low = mid
		}
	}
	return high
}

// its just a fake
func isBadVersion(n, bad int) bool {
	if n >= bad {
		return true
	}
	return false
}
