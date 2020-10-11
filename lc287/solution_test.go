package solution

import (
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	nums1 := []int{1, 3, 4, 2, 2}
	r1 := 2
	nums2 := []int{3, 1, 3, 4, 2}
	r2 := 3
	nums3 := []int{1, 1, 2}
	r3 := 1

	rsu1 := findDuplicate(nums1)
	rsu2 := findDuplicate(nums2)
	rsu3 := findDuplicate(nums3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
