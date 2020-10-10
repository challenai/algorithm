package solution

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := hash[s[i]]; ok {
			hash[s[i]]++
		} else {
			hash[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		if _, ok := hash[t[i]]; ok {
			hash[s[i]]--
			if hash[s[i]] == 0 {
				delete(hash, s[i])
			}
		} else {
			return false
		}
	}
	return true
}
