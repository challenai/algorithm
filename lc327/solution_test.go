package solution

import (
	"testing"
)

func TestCountRangeSum(t *testing.T) {
	nums1 := []int{-2, 5, -1}
	lower1 := -2
	upper1 := 2
	r1 := 3

	rsu1 := countRangeSum(nums1, lower1, upper1)

	if r1 != rsu1 {
		t.Fail()
	}
}
