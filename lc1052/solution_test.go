package solution

import (
	"testing"
)

func TestMaxSatisfied(t *testing.T) {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumps := []int{0, 1, 0, 1, 0, 1, 0, 1}
	x := 3
	r := 16

	rsu := maxSatisfied(customers, grumps, x)

	if rsu != r {
		t.Fail()
	}
}
