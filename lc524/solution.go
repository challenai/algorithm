package solution

func findLongestWord(word string, wordList []string) string {
	rsu := ""
	for i := 0; i < len(wordList); i++ {
		if isSubseq(wordList[i], word) {
			if len(wordList[i]) > len(rsu) {
				rsu = wordList[i]
				continue
			}
			if len(wordList[i]) == len(rsu) && wordList[i] < rsu {
				rsu = wordList[i]
			}
		}
	}
	return rsu
}

func isSubseq(sub, word string) bool {
	idx := 0
	for i := 0; i < len(sub); i++ {
		for idx < len(word) {
			if sub[i] == word[idx] {
				break
			}
			idx++
			if len(word) == idx {
				return false
			}
		}
	}
	return true
}
