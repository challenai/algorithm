package solution

import (
	"testing"
)

func TestNumSquares(t *testing.T) {
	n1 := 12
	r1 := 3
	n2 := 13
	r2 := 2

	resu1 := numSquares(n1)
	resu2 := numSquares(n2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
