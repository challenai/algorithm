package solution

import (
	"testing"
)

func TestIsPerfectSquare(t *testing.T) {
	num1 := 168921
	r1 := true
	num2 := 1025
	r2 := false

	rsu1 := isPerfectSquare(num1)
	rsu2 := isPerfectSquare(num2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
