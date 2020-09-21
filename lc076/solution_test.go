package solution

import (
	"testing"
)

func TestMinWindow(t *testing.T) {

	s := "ADOBECODEBANC"
	tt := "ABC"
	r := "BANC"
	resu := minWindow(s, tt)
	if resu != r {
		t.Fail()
	}
}
