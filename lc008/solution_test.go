package solution

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {
	s1 := "42 3 "
	s2 := "   -42  "
	s3 := "-123456+"
	resu1 := myAtoi(s1)
	resu2 := myAtoi(s2)
	resu3 := myAtoi(s3)

	if resu1 != 42 || resu2 != -42 || resu3 != -123456 {
		t.Fail()
	}
}
