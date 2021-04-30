package solution

func detectCapitalUse(word string) bool {
	isCapitalSeq := false
	if len(word) <= 1 {
		return true
	}

	if !isCapital(word[0]) && isCapital(word[1]) {
		return false
	}
	if isCapital(word[1]) {
		isCapitalSeq = true
	}
	for i := 2; i < len(word); i++ {
		if isCapital(word[i]) != isCapitalSeq {
			return false
		}
	}
	return true
}

func isCapital(b byte) bool {
	if b >= 65 && b <= 90 {
		return true
	}
	return false
}
