package solution

import (
	"testing"
)

func TestMaximumGap(t *testing.T) {
	nums1 := []int{3, 6, 9, 1}
	r1 := 3
	nums2 := []int{10}
	r2 := 0

	resu1 := maximumGap(nums1)
	resu2 := maximumGap(nums2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
