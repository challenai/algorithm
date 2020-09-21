package solution

import (
	"testing"
)

func TestSortColors(t *testing.T) {

	nums1 := []int{2, 0, 2, 1, 1, 0}
	r1 := []int{0, 0, 1, 1, 2, 2}
	nums2 := []int{2, 0, 1}
	r2 := []int{0, 1, 2}
	nums3 := []int{0}
	r3 := []int{0}
	sortColors(nums1)
	sortColors(nums2)
	sortColors(nums3)

	if nums1[0] != r1[0] || nums2[0] != r2[0] || nums3[0] != r3[0] {
		t.Fail()
	}
}
