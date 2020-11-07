package solution

import (
	"testing"
)

func TestRepeatedSubstringPattern(t *testing.T) {
	s1 := "abab"
	r1 := true
	s2 := "aba"
	r2 := false
	s3 := "abcabcabcabc"
	r3 := true

	rsu1 := repeatedSubstringPattern(s1)
	rsu2 := repeatedSubstringPattern(s2)
	rsu3 := repeatedSubstringPattern(s3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
