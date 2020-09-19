package solution

func lengthOfLastWord(s string) int {
	current := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			current++
		} else {
			current = 0
		}
	}
	return current
}
