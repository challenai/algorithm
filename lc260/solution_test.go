package solution

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums1 := []int{1, 2, 1, 3, 2, 5}
	r1 := []int{3, 5}
	nums2 := []int{-1, 0}
	r2 := []int{-1, 0}
	nums3 := []int{0, 1}
	r3 := []int{0, 1}

	rsu1 := singleNumber(nums1)
	rsu2 := singleNumber(nums2)
	rsu3 := singleNumber(nums3)

	if rsu1[0] != r1[0] || rsu2[1] != r2[1] || rsu3[1] != r3[1] {
		t.Fail()
	}
}
