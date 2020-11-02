package solution

import (
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	nums1 := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	r1 := 1
	nums2 := [][]int{{1, 2}, {1, 2}, {1, 2}}
	r2 := 2
	nums3 := [][]int{{1, 2}, {2, 3}}
	r3 := 0

	resu1 := eraseOverlapIntervals(nums1)
	resu2 := eraseOverlapIntervals(nums2)
	resu3 := eraseOverlapIntervals(nums3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
