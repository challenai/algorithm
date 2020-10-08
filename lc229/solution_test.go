package solution

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	nums1 := []int{3, 2, 3}
	r1 := []int{3}
	nums2 := []int{1}
	r2 := []int{1}
	nums3 := []int{1, 2}
	r3 := []int{1, 2}

	resu1 := majorityElement(nums1)
	resu2 := majorityElement(nums2)
	resu3 := majorityElement(nums3)

	if resu1[0] != r1[0] || resu2[0] != r2[0] || len(resu3) != len(r3) {
		t.Fail()
	}
}
