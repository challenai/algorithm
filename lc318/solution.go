package solution

func maxProduct(words []string) int {
	resu := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if !checkIntersect(words[i], words[j]) {
				if resu < len(words[i])*len(words[j]) {
					resu = len(words[i]) * len(words[j])
				}
			}
		}
	}
	return resu
}

func checkIntersect(word1, word2 string) bool {
	hash := map[byte]bool{}
	for i := 0; i < len(word1); i++ {
		hash[word1[i]] = false
	}
	for i := 0; i < len(word2); i++ {
		if _, ok := hash[word2[i]]; ok {
			return true
		}
	}
	return false
}
