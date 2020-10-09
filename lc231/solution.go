package solution

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n <= 2 {
		return true
	}
	counter := 2
	current := 2
	for counter > 1 {
		if current*counter == n {
			return true
		}
		if current*counter < n {
			current *= counter
			counter <<= 1
		} else {
			counter >>= 1
		}
	}
	return current == n
}
