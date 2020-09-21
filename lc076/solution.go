package solution

func minWindow(s, t string) string {
	// parse every single chars into a hashmap
	// everytime we handle a char, we try to compare wheather current string is correct
	// if correct, try to remove chars from left side.
	resu := s
	targetHash := map[byte]bool{}
	currentHash := map[byte]int{}
	for i := 0; i < len(t); i++ {
		targetHash[t[i]] = false
	}
	left := 0
	haveMatch := false
	for i := 0; i < len(s); i++ {
		if _, ok := targetHash[s[i]]; ok {
			currentHash[s[i]] = i
		}
		for len(currentHash) == len(targetHash) {
			haveMatch = true
			if i-left+1 < len(resu) {
				resu = s[left : i+1]
			}
			if idx, ok := currentHash[s[left]]; ok && idx == left {
				delete(currentHash, s[idx])
			}
			left++
		}
	}
	if !haveMatch {
		return ""
	}
	return resu
}
