package solution

import (
	"testing"
)

func TestRob(t *testing.T) {
	nums1 := []int{2, 3, 2}
	r1 := 3
	nums2 := []int{1, 2, 3, 1}
	r2 := 4
	nums3 := []int{0}
	r3 := 0

	rsu1 := rob(nums1)
	rsu2 := rob(nums2)
	rsu3 := rob(nums3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
