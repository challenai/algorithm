package solution

import (
	"testing"
)

func TestLastStoneWeight(t *testing.T) {
	nums := []int{2, 7, 4, 1, 8, 1}
	r := 1

	rsu := lastStoneWeight(nums)

	if rsu != r {
		t.Fail()
	}
}
