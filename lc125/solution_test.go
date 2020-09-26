package solution

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	s1 := "A man, a plan, a canal: Panama"
	r1 := true
	s2 := "race a car"
	r2 := false
	resu1 := isPalindrome(s1)
	resu2 := isPalindrome(s2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
