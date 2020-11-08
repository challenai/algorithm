package solution

import (
	"testing"
)

func TestTotalHammingDistance(t *testing.T) {
	nums := []int{4, 14, 2}
	r := 6

	rsu := totalHammingDistance(nums)

	if rsu != r {
		t.Fail()
	}
}
