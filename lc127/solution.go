package solution

var resu int
var existedWord map[string]bool
var transformableHash map[string][]string
var current int
var endWord_ string

func ladderLength(beginWord, endWord string, wordList []string) int {
	resu = 0
	existedWord = map[string]bool{}
	transformableHash = map[string][]string{}
	transformableHash[beginWord] = []string{}
	endWord_ = endWord
	for i := 0; i < len(wordList); i++ {
		for j := 0; j < len(wordList); j++ {
			if i == j {
				continue
			}
			if verifyTransformable(wordList[i], wordList[j]) {
				transformableHash[wordList[i]] = append(transformableHash[wordList[i]], wordList[j])
			}
		}
		if verifyTransformable(beginWord, wordList[i]) {
			transformableHash[beginWord] = append(transformableHash[beginWord], wordList[i])
		}
	}

	current = 1
	existedWord[beginWord] = false
	loop(beginWord)
	return resu
}

func verifyTransformable(a, b string) bool {
	if a == b || len(a) != len(b) {
		return false
	}
	counter := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			counter++
		}
	}
	return counter == 1
}

func loop(word string) {
	if word == endWord_ {
		if resu == 0 {
			resu = current
		} else if current < resu {
			resu = current
		}
		return
	}
	if resu != 0 && current > resu {
		return
	}

	nextSet, _ := transformableHash[word]
	for i := 0; i < len(nextSet); i++ {
		if _, ok := existedWord[nextSet[i]]; !ok {
			current++
			existedWord[nextSet[i]] = false
			loop(nextSet[i])
			current--
			delete(existedWord, nextSet[i])
		}
	}
}
