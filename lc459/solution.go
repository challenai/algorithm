package solution

func repeatedSubstringPattern(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	if len(s) == 0 {
		return true
	}
	// we can observe that if a string can be split to same substring,
	// it must can be split in the mid, even though it consist of more than 2 parts
	// what we need to care is just whether we can split it into 2 parts
	mid := len(s) >> 1
	for i := 0; i < mid; i++ {
		if s[i] != s[mid+i] {
			return false
		}
	}
	return true
}
