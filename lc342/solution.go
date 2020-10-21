package solution

func isPowerOfFour(num int) bool {
	if num < 3 {
		return false
	}
	for num%4 == 0 {
		num /= 4
	}
	return num == 1
}
