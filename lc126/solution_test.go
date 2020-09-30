package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	beginWord1 := "hit"
	endWord1 := "cog"
	wordList1 := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	beginWord2 := "hit"
	endWord2 := "cog"
	wordList2 := []string{"hot", "dot", "dog", "lot", "log"}
	r1 := [][]string{
		{"hit", "hot", "dot", "dog", "cog"},
		{"hit", "hot", "lot", "log", "cog"},
	}
	r2 := [][]string{}
	resu1 := findLadders(beginWord1, endWord1, wordList1)
	resu2 := findLadders(beginWord2, endWord2, wordList2)

	if len(resu1) != len(r1) || len(resu2) != len(r2) {
		t.Fail()
	}
}
