package solution

import (
	"testing"
)

func TestAddDigits(t *testing.T) {
	n := 38
	r := 2

	resu := addDigits(n)

	if resu != r {
		t.Fail()
	}
}
