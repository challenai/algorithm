package solution

import (
	"testing"
)

func TestReverseBits(t *testing.T) {
	num1 := 43261596
	r1 := 964176192
	num2 := 4294967293
	r2 := 3221225471

	resu1 := reverseBits(num1)
	resu2 := reverseBits(num2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
