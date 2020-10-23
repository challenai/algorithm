package solution

import (
	"testing"
)

func TestLargestDivisibleSubset(t *testing.T) {
	nums1 := []int{1, 2, 3}
	r1 := []int{1, 2}
	nums2 := []int{1, 2, 4, 8}
	r2 := []int{1, 2, 4, 8}
	nums3 := []int{2, 3, 4, 5, 6, 9, 18}
	r3 := []int{3, 9, 18}

	rsu1 := largestDivisibleSubset(nums1)
	rsu2 := largestDivisibleSubset(nums2)
	rsu3 := largestDivisibleSubset(nums3)

	if len(r1) != len(rsu1) || len(r2) != len(rsu2) || len(r3) != len(rsu3) {
		t.Fail()
	}
}
