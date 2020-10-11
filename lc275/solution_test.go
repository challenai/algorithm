package solution

import (
	"testing"
)

func TestHIndex2(t *testing.T) {
	citations := []int{0, 1, 3, 5, 6}
	r := 3
	resu := hIndex2(citations)

	if resu != r {
		t.Fail()
	}
}
