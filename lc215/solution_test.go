package solution

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	nums1 := []int{3, 2, 1, 5, 6, 4}
	k1 := 2
	r1 := 5
	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k2 := 4
	r2 := 4

	rsu1 := findKthLargest(nums1, k1)
	rsu2 := findKthLargest(nums2, k2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
