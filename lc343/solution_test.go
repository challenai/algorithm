package solution

import (
	"testing"
)

func TestIntegerBreak(t *testing.T) {
	n1 := 2
	r1 := 1
	n2 := 10
	r2 := 36
	rsu1 := integerBreak(n1)
	rsu2 := integerBreak(n2)

	if r1 != rsu1 || r2 != rsu2 {
		t.Fail()
	}
}
