package solution

import (
	"testing"
)

func TestReverse(t *testing.T) {
	x1 := 123
	r1 := 321
	x2 := 120
	r2 := 21

	resu1 := reverse(x1)
	resu2 := reverse(x2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
