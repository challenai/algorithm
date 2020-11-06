package solution

import (
	"testing"
)

func TestFind132pattern(t *testing.T) {
	nums1 := []int{1, 2, 3, 4}
	r1 := false
	nums2 := []int{3, 1, 4, 2}
	r2 := true
	nums3 := []int{-1, 3, 2, 0}
	r3 := true

	rsu1 := find132pattern(nums1)
	rsu2 := find132pattern(nums2)
	rsu3 := find132pattern(nums3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
