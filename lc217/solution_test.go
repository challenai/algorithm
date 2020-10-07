package solution

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	nums1 := []int{1, 2, 3, 1}
	r1 := true
	nums2 := []int{1, 2, 3, 4}
	r2 := false
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	r3 := true

	resu1 := containsDuplicate(nums1)
	resu2 := containsDuplicate(nums2)
	resu3 := containsDuplicate(nums3)

	if r1 != resu1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
