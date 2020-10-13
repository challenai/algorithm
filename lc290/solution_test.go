package solution

import (
	"testing"
)

func TestWordPattern(t *testing.T) {
	pattern1 := "abba"
	s1 := "dog cat cat dog"
	r1 := true
	pattern2 := "abba"
	s2 := "dog cat cat fish"
	r2 := false
	pattern3 := "aaaa"
	s3 := "dog cat cat dog"
	r3 := false
	pattern4 := "abba"
	s4 := "dog dog dog dog"
	r4 := false

	resu1 := wordPattern(pattern1, s1)
	resu2 := wordPattern(pattern2, s2)
	resu3 := wordPattern(pattern3, s3)
	resu4 := wordPattern(pattern4, s4)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 || resu4 != r4 {
		t.Fail()
	}
}
