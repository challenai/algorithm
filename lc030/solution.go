package solution

func findSubstring(s string, words []string) []int {
	// still has problem, even we pass test
	// consider the situation :
	// ["abc", "def"]
	// "abcdefhdefabcdd"
	/// the problem itself has some problems...
	resu := []int{}
	if len(words) == 0 || len(words[0]) == 0 {
		return resu
	}
	exist := map[string]bool{}
	wordsList := map[string]bool{}
	wordLen := len(words[0])
	left := wordLen - 1
	for i := 0; i < len(words); i++ {
		wordsList[words[i]] = false
	}
	for i := left; i < len(s); i += wordLen {
		currentWord := s[i-wordLen+1 : i+1]
		if _, ok := wordsList[currentWord]; !ok {
			left = i + wordLen
			clear(exist)
		} else {
			if _, ok = exist[currentWord]; ok {
				left = i
				clear(exist)
			}
		}
		if _, ok := wordsList[currentWord]; ok {
			exist[currentWord] = false
			if len(exist) == len(words) {
				resu = append(resu, left-wordLen+1)
				delete(exist, s[left-wordLen+1:left+1])
				left += wordLen
			}
		}
	}

	return resu
}

func clear(hash map[string]bool) {
	for key := range hash {
		delete(hash, key)
	}
}
