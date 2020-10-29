package solution

import (
	"testing"
)

func TestSplitArray(t *testing.T) {
	nums1 := []int{7, 2, 5, 10, 8}
	m1 := 2
	r1 := 18
	nums2 := []int{1, 2, 3, 4, 5}
	m2 := 2
	r2 := 9
	nums3 := []int{1, 4, 4}
	m3 := 3
	r3 := 4
	rsu1 := splitArray(nums1, m1)
	rsu2 := splitArray(nums2, m2)
	rsu3 := splitArray(nums3, m3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
