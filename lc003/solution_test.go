package solution

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s1 := ""
	r1 := 0
	s2 := "bbbbb"
	r2 := 1
	s3 := "pwwkew"
	r3 := 3
	resu1 := lengthOfLongestSubstring(s1)
	resu2 := lengthOfLongestSubstring(s2)
	resu3 := lengthOfLongestSubstring(s3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
