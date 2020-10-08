package solution

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	s1 := "1 + 143"
	r1 := 144
	s2 := " 2-1 2+ 2 "
	r2 := -8
	s3 := "(1+(4+5+2)-3)+(6+8)"
	r3 := 23

	rsu1 := calculate(s1)
	rsu2 := calculate(s2)
	rsu3 := calculate(s3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
