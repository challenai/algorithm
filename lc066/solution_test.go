package solution

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	digits1 := []int{1, 2, 3}
	r1 := []int{1, 2, 4}

	digits2 := []int{4, 3, 2, 1}
	r2 := []int{4, 3, 2, 2}

	digits3 := []int{0}
	r3 := []int{1}

	resu1 := plusOne(digits1)
	resu2 := plusOne(digits2)
	resu3 := plusOne(digits3)

	if resu1[2] != r1[2] || resu2[3] != r2[3] || resu3[0] != r3[0] {
		t.Fail()
	}
}
