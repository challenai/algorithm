package solution

func isUgly(n int) bool {
	for n%5 == 0 {
		n /= 5
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%2 == 0 {
		n /= 2
	}
	return n == 1
}
