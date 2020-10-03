package solution

import (
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	n1 := 3
	r1 := 0
	n2 := 5
	r2 := 1
	n3 := 0
	r3 := 0

	resu1 := trailingZeroes(n1)
	resu2 := trailingZeroes(n2)
	resu3 := trailingZeroes(n3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
