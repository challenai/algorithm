package solution

import (
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	num1 := 1
	r1 := "A"
	num2 := 28
	r2 := "AB"
	num3 := 701
	r3 := "ZY"

	resu1 := convertToTitle(num1)
	resu2 := convertToTitle(num2)
	resu3 := convertToTitle(num3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
