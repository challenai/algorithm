package solution

func longestPalindrome(s string) int {
	resu := 0
	hash := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		if _, ok := hash[s[i]]; ok {
			resu += 2
			delete(hash, s[i])
			continue
		}
		hash[s[i]] = false
	}
	if len(hash) > 0 {
		resu++
	}
	return resu
}
