package solution

import (
	"testing"
)

func TestIsPowerOfFour(t *testing.T) {
	num1 := 16
	r1 := true
	num2 := 5
	r2 := false

	rsu1 := isPowerOfFour(num1)
	rsu2 := isPowerOfFour(num2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
