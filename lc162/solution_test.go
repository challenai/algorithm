package solution

import (
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 1, 3, 5, 6, 4}

	resu1 := findPeakElement(nums1)
	resu2 := findPeakElement(nums2)

	if resu1 != 2 || (resu2 != 1 && resu2 != 5) {
		t.Fail()
	}
}
