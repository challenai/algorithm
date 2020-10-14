package solution

import (
	"testing"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	s1 := "bcabc"
	r1 := "abc"
	s2 := "cbacdcbc"
	r2 := "acdb"

	rsu1 := removeDuplicateLetters(s1)
	rsu2 := removeDuplicateLetters(s2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
