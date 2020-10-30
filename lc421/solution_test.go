package solution

import (
	"testing"
)

func TestFindMaximumXOR(t *testing.T) {
	nums1 := []int{3, 10, 5, 25, 2, 8}
	r1 := 28
	nums2 := []int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}
	r2 := 127
	nums3 := []int{0}
	r3 := 0

	rsu1 := findMaximumXOR(nums1)
	rsu2 := findMaximumXOR(nums2)
	rsu3 := findMaximumXOR(nums3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
