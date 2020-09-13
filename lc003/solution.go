package solution

func lengthOfLongestSubstring(s string) int {
	// idea: use a window to store no dulplicate data, but we dont necessarly delete element in window
	if len(s) == 0 {
		return 0
	}
	hash := map[byte]int{}
	beforePtr := 0
	currentMaxLen := 0
	for i := 0; i < len(s); i++ {
		if idx, ok := hash[s[i]]; ok && beforePtr <= idx {
			beforePtr = idx + 1
		} else {
			currentMaxLen = max(currentMaxLen, i-beforePtr+1)
		}
		hash[s[i]] = i
	}

	return currentMaxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
