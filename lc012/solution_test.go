package solution

import (
	"testing"
)

func TestIntToRoman(t *testing.T) {
	num1 := 3
	r1 := "III"
	num2 := 4
	r2 := "IV"
	num3 := 9
	r3 := "IX"
	num4 := 58
	r4 := "LVIII"
	num5 := 1994
	r5 := "MCMXCIV"

	resu1 := intToRoman(num1)
	resu2 := intToRoman(num2)
	resu3 := intToRoman(num3)
	resu4 := intToRoman(num4)
	resu5 := intToRoman(num5)

	if resu1 != r1 && resu2 != r2 && resu3 != r3 && resu4 != r4 && resu5 != r5 {
		t.Fail()
	}
	if false {
		t.Fail()
	}
}
