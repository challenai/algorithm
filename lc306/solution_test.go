package solution

import (
	"testing"
)

func TestIsAdditiveNumber(t *testing.T) {
	num1 := "112358"
	r1 := true
	num2 := "199100199"
	r2 := true
	num3 := "114342"
	r3 := false

	rsu1 := isAdditiveNumber(num1)
	rsu2 := isAdditiveNumber(num2)
	rsu3 := isAdditiveNumber(num3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
