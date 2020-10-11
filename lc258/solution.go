package solution

func addDigits(n int) int {
	if n <= 9 {
		return n
	}
	current := 0
	for n > 9 {
		for n > 0 {
			current += n % 10
			n /= 10
		}
		n = current
		current = 0
	}
	return n
}
