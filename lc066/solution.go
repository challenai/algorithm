package solution

func plusOne(digits []int) []int {
	allNine := true
	for i := 0; i < len(digits); i++ {
		if digits[i] != 9 {
			allNine = false
		}
	}
	if allNine {
		digits = append(digits, 0)
		digits[0] = 1
		for i := 1; i < len(digits); i++ {
			digits[i] = 0
		}
		return digits
	}

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] > 9 {
			digits[i] -= 10
		} else {
			break
		}
	}
	return digits
}
