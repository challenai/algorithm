package solution

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3, 4, 5, 16}
	nums2 := []int{2, 8, 15, 17}
	resu := findMedianSortedArrays(nums1, nums2)

	if resu != 5 {
		t.Fail()
	}
}
