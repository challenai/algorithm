package solution

import (
	"testing"
)

func TestLongestPalindrome3(t *testing.T) {
	str1 := "babad"
	r1 := "bab"
	str2 := "cbbd"
	r2 := "bb"
	resu1 := longestPalindrome(str1)
	resu2 := longestPalindrome(str2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
