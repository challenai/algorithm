package solution

import (
	"testing"
)

func TestPalindromePairs(t *testing.T) {
	words0 := []string{"abcd", "dcba", "lls", "s", "sssll"}
	words1 := []string{"bat", "tab", "cat"}
	words2 := []string{"a", ""}
	r1 := [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}}
	r2 := [][]int{{0, 1}, {1, 0}}
	r3 := [][]int{{0, 1}, {1, 0}}

	rsu1 := palindromePairs(words0)
	rsu2 := palindromePairs(words1)
	rsu3 := palindromePairs(words2)

	if len(r1) != len(rsu1) || len(r2) != len(rsu2) || len(r3) != len(rsu3) {
		t.Fail()
	}
}
