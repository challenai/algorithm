package solution

import (
	"testing"
)

func TestFindComplement(t *testing.T) {
	num1 := 5
	r1 := 2
	num2 := 1
	r2 := 0
	num3 := 53852347
	r3 := 9144652

	rsu1 := findComplement(num1)
	rsu2 := findComplement(num2)
	rsu3 := findComplement(num3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
