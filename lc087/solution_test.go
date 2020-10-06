package solution

import (
	"testing"
)

func TestIsScramble(t *testing.T) {
	s11 := "great"
	s12 := "rgeat"
	r1 := true
	s21 := "abcde"
	s22 := "caebd"
	r2 := false
	s31 := "a"
	s32 := "a"
	r3 := true

	rsu1 := isScramble(s11, s12)
	rsu2 := isScramble(s21, s22)
	rsu3 := isScramble(s31, s32)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
