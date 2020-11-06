package solution

import (
	"testing"
)

func TestMinMoves2(t *testing.T) {
	nums1 := []int{1, 2, 3}
	r1 := 2
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	r2 := 36

	rsu1 := minMoves2(nums1)
	rsu2 := minMoves2(nums2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
