package solution

import (
	"testing"
)

func TestFindKthNumber(t *testing.T) {
	n := 13
	k := 2
	r := 10

	rsu := findKthNumber(n, k)

	if rsu != r {
		t.Fail()
	}
}
