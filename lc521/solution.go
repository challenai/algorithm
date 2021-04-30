package solution

func findLUSlength(a, b string) int {
	if a == b {
		return -1
	}
	if len(a) != len(b) {
		return max(len(a), len(b))
	}
	return len(a)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
