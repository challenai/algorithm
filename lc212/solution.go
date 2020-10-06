package solution

import (
	"strconv"
)

var path map[string]bool
var resu []string
var board_ [][]byte
var currentWord string
var currentIdx int

func findWords(board [][]byte, words []string) []string {
	board_ = board
	resu = []string{}
	for i := 0; i < len(words); i++ {
		currentWord = words[i]
		existWord()
	}

	return resu
}

func delveIntoNext(i, j int) {
	if board_[i][j] == currentWord[currentIdx] {
		if _, ok := path[serialize(i, j)]; !ok {
			path[serialize(i, j)] = false
			currentIdx++
			dfs(i, j)
			delete(path, serialize(i, j))
			currentIdx--
		}
	}
}

func dfs(i, j int) {
	if currentIdx >= len(currentWord)-1 {
		resu = append(resu, currentWord)
		return
	}
	if i > 0 {
		delveIntoNext(i-1, j)
	}
	if i < len(board_)-1 {
		delveIntoNext(i+1, j)
	}
	if j > 0 {
		delveIntoNext(i, j-1)
	}
	if j < len(board_[0])-1 {
		delveIntoNext(i, j+1)
	}
}

func existWord() bool {
	for i := 0; i < len(board_); i++ {
		for j := 0; j < len(board_[i]); j++ {
			currentIdx = 0
			if board_[i][j] == currentWord[currentIdx] {
				currentIdx++
				path = map[string]bool{
					serialize(i, j): false,
				}
				dfs(i, j)
			}
		}
	}
	return false
}

func serialize(i, j int) string {
	return strconv.Itoa(i) + "_" + strconv.Itoa(j)
}
