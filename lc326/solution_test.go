package solution

import (
	"testing"
)

func TestIsPowerOfThree(t *testing.T) {
	n1 := 27
	r1 := true
	n2 := 0
	r2 := false
	n3 := 9
	r3 := true
	n4 := 45
	r4 := false

	rsu1 := isPowerOfThree(n1)
	rsu2 := isPowerOfThree(n2)
	rsu3 := isPowerOfThree(n3)
	rsu4 := isPowerOfThree(n4)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 || rsu4 != r4 {
		t.Fail()
	}
}
