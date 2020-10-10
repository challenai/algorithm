package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	num1 := 6
	r1 := true
	num2 := 8
	r2 := true
	num3 := 14
	r3 := false

	resu1 := isUgly(num1)
	resu2 := isUgly(num2)
	resu3 := isUgly(num3)

	if resu3 != r3 || resu2 != r2 || resu1 != r1 {
		t.Fail()
	}
}
