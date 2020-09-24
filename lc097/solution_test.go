package solution

import (
	"testing"
)

func TestIsInterleave(t *testing.T) {
	s11 := "aabcc"
	s12 := "dbbca"
	s13 := "aadbbcbcac"
	r1 := true
	s21 := "aabcc"
	s22 := "dbbca"
	s23 := "aadbbbaccc"
	r2 := false

	resu1 := isInterleave(s11, s12, s13)
	resu2 := isInterleave(s21, s22, s23)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
