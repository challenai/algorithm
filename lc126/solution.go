package solution

var resu [][]string
var current []string
var existedPath map[string]bool
var transformHash map[string][]string
var endWord_ string

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	resu = [][]string{}
	current = []string{}
	existedPath = map[string]bool{}
	endWord_ = endWord
	transformHash = map[string][]string{}
	transformHash[beginWord] = []string{}
	transformHash[endWord] = []string{}
	for i := 0; i < len(wordList); i++ {
		transformHash[wordList[i]] = []string{}
		for j := 0; j < len(wordList); j++ {
			if i == j {
				continue
			}
			if verifyTransformable(wordList[i], wordList[j]) {
				transformHash[wordList[i]] = append(transformHash[wordList[i]], wordList[j])
			}
		}
		if verifyTransformable(wordList[i], beginWord) {
			transformHash[beginWord] = append(transformHash[beginWord], wordList[i])
		}
		if verifyTransformable(wordList[i], endWord) {
			transformHash[endWord] = append(transformHash[endWord], wordList[i])
		}
	}
	current = append(current, beginWord)
	existedPath[beginWord] = false
	loop(beginWord)

	minLen := len(wordList) + 2
	for i := 0; i < len(resu); i++ {
		if minLen > len(resu[i]) {
			minLen = len(resu[i])
		}
	}
	resuFiltered := [][]string{}
	for i := 0; i < len(wordList); i++ {
		if len(wordList) == minLen {
			resuFiltered = append(resuFiltered, wordList[i])
		}
	}

	return resuFiltered
}

func verifyTransformable(a, b string) bool {
	counter := 0
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			counter++
		}
	}
	return counter == 1
}

func loop(currentWord string) {
	if currentWord == endWord_ {
		temp := make([]string, len(current))
		copy(temp, current)
		resu = append(resu, temp)
		return
	}
	transformableList, _ := transformHash[currentWord]
	for i := 0; i < len(transformableList); i++ {
		if _, ok := existedPath[transformableList[i]]; !ok {
			current = append(current, transformableList[i])
			existedPath[transformableList[i]] = false
			loop(transformableList[i])
			current = current[:len(current)-1]
			delete(existedPath, transformableList[i])
		}
	}
}
