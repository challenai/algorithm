package solution

import (
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums1 := []int{5, 7, 7, 8, 8, 10}
	target1 := 8
	r1 := []int{3, 4}
	nums2 := []int{5, 7, 7, 8, 8, 10}
	target2 := 6
	r2 := []int{-1, -1}

	resu1 := searchRange(nums1, target1)
	resu2 := searchRange(nums2, target2)

	if resu1[0] != r1[0] || resu2[1] != r2[1] {
		t.Fail()
	}
}
