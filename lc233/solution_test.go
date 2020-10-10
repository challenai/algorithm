package solution

import (
	"testing"
)

func TestCountDigitOne(t *testing.T) {
	// need more test case
	n1 := 13
	r1 := 6

	resu1 := countDigitOne(n1)

	if resu1 != r1 {
		t.Fail()
	}
}
