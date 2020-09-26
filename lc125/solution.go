package solution

func isPalindrome(s string) bool {
	var left, right int
	left = 0
	right = len(s) - 1
	for left < right {
		for !checkAlpha(s[left]) {
			left++
		}
		for !checkAlpha(s[right]) {
			right--
		}
		if left < right {
			if !alphaCompare(s[left], s[right]) {
				return false
			}
		}
		left++
		right--
	}
	return true
}

func checkAlpha(bt byte) bool {
	if (bt <= 'z' && bt >= 'a') || (bt <= 'Z' && bt >= 'A') {
		return true
	}
	return false
}

func alphaCompare(bt1, bt2 byte) bool {
	if bt1 == bt2 {
		return true
	}
	if bt1 < 'a' {
		// haha, I learn something here, 'a'-'A' = 32, wagada! remember that 'a' is bigger than 'A'...
		bt1 += 32
	}
	if bt2 < 'a' {
		bt2 += 32
	}
	return bt1 == bt2
}
