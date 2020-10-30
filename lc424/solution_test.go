package solution

import (
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	s1 := "ABAB"
	k1 := 2
	r1 := 4
	s2 := "AABABBA"
	k2 := 1
	r2 := 4

	rsu1 := characterReplacement(s1, k1)
	rsu2 := characterReplacement(s2, k2)

	if r1 != rsu1 || rsu2 != r2 {
		t.Fail()
	}
}
