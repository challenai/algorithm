package solution

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	s1 := "aa"
	p1 := "a"
	r1 := false

	s2 := "aa"
	p2 := "a*"
	r2 := true

	s3 := "ab"
	p3 := ".*"
	r3 := true

	s4 := "aab"
	p4 := "c*a*b"
	r4 := true

	s5 := "mississippi"
	p5 := "mis*is*p*."
	r5 := false

	resu1 := isMatch(s1, p1)
	resu2 := isMatch(s2, p2)
	resu3 := isMatch(s3, p3)
	resu4 := isMatch(s4, p4)
	resu5 := isMatch(s5, p5)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 || resu4 != r4 || resu5 != r5 {
		t.Fail()
	}

}
