package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	n1 := 5
	b1 := 4
	r1 := 4
	n2 := 1
	b2 := 1
	r2 := 1

	resu1 := firstBadVersion(n1, b1)
	resu2 := firstBadVersion(n2, b2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
