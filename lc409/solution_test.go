package solution

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	s1 := "abccccdd"
	r1 := 7
	s2 := "a"
	r2 := 1
	s3 := "bb"
	r3 := 2
	s4 := "fdhsjdWWja"
	r4 := 7

	rsu1 := longestPalindrome(s1)
	rsu2 := longestPalindrome(s2)
	rsu3 := longestPalindrome(s3)
	rsu4 := longestPalindrome(s4)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 || rsu4 != r4 {
		t.Fail()
	}
}
