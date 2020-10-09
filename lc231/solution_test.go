package solution

import (
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	n1 := 1
	r1 := true
	n2 := 16
	r2 := true
	n3 := 3
	r3 := false
	n4 := 2048
	r4 := true
	n5 := 1025
	r5 := false

	resu1 := isPowerOfTwo(n1)
	resu2 := isPowerOfTwo(n2)
	resu3 := isPowerOfTwo(n3)
	resu4 := isPowerOfTwo(n4)
	resu5 := isPowerOfTwo(n5)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 || resu4 != r4 || resu5 != r5 {
		t.Fail()
	}
}
