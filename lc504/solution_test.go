package solution

import (
	"testing"
)

func TestConvertToBase7(t *testing.T) {

	num1 := 100
	r1 := "202"
	num2 := -7
	r2 := "-10"
	rsu1 := convertToBase7(num1)
	rsu2 := convertToBase7(num2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
