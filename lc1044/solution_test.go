package solution

import (
	"testing"
)

func TestLongestDupSubstring(t *testing.T) {
	s1 := "banana"
	r1 := "ana"
	s2 := "abcd"
	r2 := ""

	rsu1 := longestDupSubstring(s1)
	rsu2 := longestDupSubstring(s2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
