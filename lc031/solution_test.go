package solution

import (
	"testing"
)

func TestNextPermutation(t *testing.T) {
	nums1 := []int{1, 2, 3}
	r1 := []int{1, 3, 2}
	nums2 := []int{3, 2, 1}
	r2 := []int{1, 2, 3}
	nums3 := []int{1, 1, 5}
	r3 := []int{1, 5, 1}

	nextPermutation(nums1)
	nextPermutation(nums2)
	nextPermutation(nums3)

	pass := true
	if !compareSlice(r1, nums1) {
		pass = false
	}

	if !compareSlice(r2, nums2) {
		pass = false
	}

	if !compareSlice(r3, nums3) {
		pass = false
	}

	if !pass {
		t.Fail()
	}
}

func compareSlice(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

