package solution

import (
	"testing"
)

func TestRotate(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	k1 := 3
	r1 := []int{5, 6, 7, 1, 2, 3, 4}
	nums2 := []int{-1, -100, 3, 99}
	k2 := 2
	r2 := []int{3, 99, -1, -100}

	rotate(nums1, k1)
	rotate(nums2, k2)

	if nums1[0] != r1[0] || nums2[2] != r2[2] {
		t.Fail()
	}
}
