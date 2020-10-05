package solution

import (
	"testing"
)

func TestRangeBitwiseAnd(t *testing.T) {
	m1 := 5
	n1 := 6
	r1 := 4
	m2 := 0
	n2 := 1
	r2 := 0

	resu1 := rangeBitwiseAnd(m1, n1)
	resu2 := rangeBitwiseAnd(m2, n2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
