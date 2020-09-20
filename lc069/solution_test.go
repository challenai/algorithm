package solution

import (
	"testing"
)

func TestSqrt(t *testing.T) {

	x1 := 4
	r1 := 2
	x2 := 8
	r2 := 2
	x3 := 2501
	r3 := 50
	resu1 := sqrt(x1)
	resu2 := sqrt(x2)
	resu3 := sqrt(x3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
