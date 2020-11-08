package solution

var rsu, words_ []string
var current string
var wordsHash map[string]string
var maxWordLen, counter int

func findAllConcatenatedWordsInADict(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}
	wordsHash = map[string]string{}
	maxWordLen = len(words[0])
	for i := 0; i < len(words); i++ {
		wordsHash[words[i]] = words[i]
		if len(words[i]) > maxWordLen {
			maxWordLen = len(words[i])
		}
	}
	words_ = words
	current = ""
	counter = 0
	dfs()
	return rsu
}

func dfs() {
	if v, ok := wordsHash[current]; ok && counter > 1 {
		rsu = append(rsu, v)
	}
	if len(current) >= maxWordLen {
		return
	}
	for i := 0; i < len(words_); i++ {
		counter++
		current += words_[i]
		dfs()
		counter--
		current = current[:len(current)-len(words_[i])]
	}
}
