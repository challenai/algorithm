package solution

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	s1 := "3+2*2"
	r1 := 7
	s2 := " 3/2 "
	r2 := 1
	s3 := " 3+5 / 2 "
	r3 := 5
	s4 := "2+4* 4/5-6 /5  +4-6 2/2 +5  *32  "
	r4 := 137

	resu1 := calculate(s1)
	resu2 := calculate(s2)
	resu3 := calculate(s3)
	resu4 := calculate(s4)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 || resu4 != r4 {
		t.Fail()
	}
}
