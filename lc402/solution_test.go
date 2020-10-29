package solution

import (
	"testing"
)

func TestRemoveKdigits(t *testing.T) {
	num1 := "1432219"
	k1 := 3
	r1 := "1219"
	num2 := "10200"
	k2 := 1
	r2 := "200"
	num3 := "10"
	k3 := 2
	r3 := "0"
	rsu1 := removeKdigits(num1, k1)
	rsu2 := removeKdigits(num2, k2)
	rsu3 := removeKdigits(num3, k3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
