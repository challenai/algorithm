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
	r1 := 5
	r2 := 0
	resu1 := ladderLength(beginWord1, endWord1, wordList1)
	resu2 := ladderLength(beginWord2, endWord2, wordList2)
	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
