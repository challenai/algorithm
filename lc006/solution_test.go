package solution

import (
	"testing"
)

func TestConvert(t *testing.T) {
	s1 := "LEETCODEISHIRING"
	n1 := 3
	r1 := "LCIRETOESIIGEDHN"

	s2 := "LEETCODEISHIRING"
	n2 := 4
	r2 := "LDREOEIIECIHNTSG"

	resu1 := convert(s1, n1)
	resu2 := convert(s2, n2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
