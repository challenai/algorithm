package solution

import (
	"testing"
)

func TestAddBinary(t *testing.T) {

	a1 := "11"
	b1 := "1"
	r1 := "100"

	a2 := "1010"
	b2 := "1011"
	r2 := "10101"

	resu1 := addBinary(a1, b1)
	resu2 := addBinary(a2, b2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
