package solution

import (
	"testing"
)

func TestCombine(t *testing.T) {
	n := 4
	k := 2
	r := 6
	resu := combine(n, k)
	if len(resu) != r {
		t.Fail()
	}
}
