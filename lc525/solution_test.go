package solution

import (
	"testing"
)

func TestFindMaxLength(t *testing.T) {
	nums1 := []int{0, 1}
	r1 := 2
	nums2 := []int{0, 1, 0}
	r2 := 2

	rsu1 := findMaxLength(nums1)
	rsu2 := findMaxLength(nums2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
