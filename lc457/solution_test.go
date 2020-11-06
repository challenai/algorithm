package solution

import (
	"testing"
)

func TestCircularArrayLoop(t *testing.T) {
	nums1 := []int{2, -1, 1, 2, 2}
	r1 := true
	nums2 := []int{-1, 2}
	r2 := false
	nums3 := []int{-2, 1, -1, -2, -2}
	r3 := false

	rsu1 := circularArrayLoop(nums1)
	rsu2 := circularArrayLoop(nums2)
	rsu3 := circularArrayLoop(nums3)

	if r1 != rsu1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
