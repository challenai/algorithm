package solution

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	num1 := 3
	r1 := 2
	num2 := 9
	r2 := 2

	resu1 := hammingWeight(num1)
	resu2 := hammingWeight(num2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
