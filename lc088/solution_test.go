package solution

import (
	"testing"
)

func TestMerge(t *testing.T) {

	nums1 := []int{1, 2, 3, 7, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 4
	n := 3
	r := []int{1, 2, 2, 3, 5, 6, 7}

	merge(nums1, m, nums2, n)
	for i := 0; i < len(r); i++ {
		if r[i] != nums1[i] {
			t.Fail()
		}
	}
}
