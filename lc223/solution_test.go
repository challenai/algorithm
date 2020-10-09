package solution

import (
	"testing"
)

func TestComputeArea(t *testing.T) {
	r := 45
	rsu := computeArea(-3, 0, 3, 4, 0, -1, 9, 2)
	if rsu != r {
		t.Fail()
	}
}
