package solution

import (
	"testing"
)

func TestIsHappy(t *testing.T) {
	num1 := 19
	r1 := true
	num2 := 1
	r2 := true

	resu1 := isHappy(num1)
	resu2 := isHappy(num2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
