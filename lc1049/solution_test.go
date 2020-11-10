package solution

import (
	"testing"
)

func TestLastStoneWeightII(t *testing.T) {
	stones := []int{2, 7, 4, 1, 8, 1}
	r := 1

	rsu := lastStoneWeightII(stones)

	if rsu != r {
		t.Fail()
	}
}
