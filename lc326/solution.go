package solution

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	endFilter := [10]bool{}
	endFilter[1] = true
	endFilter[3] = true
	endFilter[7] = true
	endFilter[9] = true
	if !endFilter[n%10] {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}
