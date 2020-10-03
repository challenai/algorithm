package solution

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	nums1 := []int{3, 2, 3}
	r1 := 3
	nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	r2 := 2

	resu1 := majorityElement(nums1)
	resu2 := majorityElement(nums2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
