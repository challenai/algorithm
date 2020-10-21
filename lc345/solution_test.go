package solution

import (
	"testing"
)

func TestReverseVowels(t *testing.T) {
	s1 := "hello"
	r1 := "holle"
	s2 := "leetcode"
	r2 := "leotcede"

	rsu1 := reverseVowels(s1)
	rsu2 := reverseVowels(s2)

	if r1 != rsu1 || r2 != rsu2 {
		t.Fail()
	}
}
