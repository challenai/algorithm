package solution

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	s1 := "III"
	r1 := 3
	s2 := "IV"
	r2 := 4
	s3 := "IX"
	r3 := 9
	s4 := "LVIII"
	r4 := 58
	resu1 := romanToInt(s1)
	resu2 := romanToInt(s2)
	resu3 := romanToInt(s3)
	resu4 := romanToInt(s4)

	if resu1 != r1 && resu2 != r2 && resu3 != r3 && resu4 != r4 {
		t.Fail()
	}
}
