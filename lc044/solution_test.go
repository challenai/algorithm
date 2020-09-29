package solution

import (
	"testing"
)

func TestX(t *testing.T) {

	s1 := "aa"
	p1 := "a"
	r1 := false
	s2 := "aa"
	p2 := "*"
	r2 := true
	s3 := "cb"
	p3 := "*a*b"
	r3 := true
	s4 := "acdcb"
	p4 := "a*c?b"
	r4 := false

	resu1 := isMatch(s1, p1)
	resu2 := isMatch(s2, p2)
	resu3 := isMatch(s3, p3)
	resu4 := isMatch(s4, p4)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 || resu4 != r4 {
		t.Fail()
	}
}
