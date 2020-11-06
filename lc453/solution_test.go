package solution

import (
	"testing"
)

func TestMinMoves(t *testing.T) {
	nums1 := []int{1, 2, 3}
	r1 := 3
	nums2 := []int{3, 4, 4, 6, 4, 5, 5, 4}
	r2 := 13
	nums3 := []int{1, 36, 7, 9, 9, 3, 6, 7}
	r3 := 210

	rsu1 := minMoves(nums1)
	rsu2 := minMoves(nums2)
	rsu3 := minMoves(nums3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
