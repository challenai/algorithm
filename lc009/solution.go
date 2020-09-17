package solution

func isPalindrome(n int) bool {
	// idea: first we try to figure out the length of this number, then use this length to set 2 pointers to move from both sides to mid.
	// test every pair if they are symentic.
	if n < 0 {
		return false
	}
	var current, counter int
	current = n
	counter = 0
	for current > 0 {
		current /= 10
		counter++
	}
	current = n
	counter--
	for current > 9 {
		if current%10 != current/(pow10(counter)) {
			return false
		}
		current %= counter * 10
		current /= 10
		counter -= 2
	}
	return true
}

func pow10(n int) int {
	if n == 0 {
		return 1
	}
	resu := 1
	for n > 0 {
		resu *= 10
		n--
	}
	return resu
}
