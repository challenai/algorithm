package solution

import (
	"testing"
)

func TestHIndex(t *testing.T) {
	citations := []int{3, 0, 6, 1, 5}
	r := 3

	resu := hIndex(citations)

	if resu != r {
		t.Fail()
	}
}
