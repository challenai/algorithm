package solution

import (
	"testing"
)

func TestCheckSubarraySum(t *testing.T) {
	nums1 := []int{23, 2, 4, 6, 7}
	k1 := 6
	r1 := true
	nums2 := []int{23, 2, 6, 4, 7}
	k2 := 6
	r2 := true
	nums3 := []int{23, 2, 6, 4, 7}
	k3 := 13
	r3 := false

	rsu1 := checkSubarraySum(nums1, k1)
	rsu2 := checkSubarraySum(nums2, k2)
	rsu3 := checkSubarraySum(nums3, k3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
