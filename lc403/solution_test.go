package solution

import (
	"testing"
)

func TestCanCross(t *testing.T) {
	nums1 := []int{0, 1, 3, 5, 6, 8, 12, 17}
	r1 := true
	nums2 := []int{0, 1, 2, 3, 4, 8, 9, 11}
	r2 := false

	rsu1 := canCross(nums1)
	rsu2 := canCross(nums2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
