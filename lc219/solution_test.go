package solution

import (
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	nums1 := []int{1, 2, 3, 1}
	k1 := 3
	r1 := true
	nums2 := []int{1, 0, 1, 1}
	k2 := 1
	r2 := true
	nums3 := []int{1, 2, 3, 1, 2, 3}
	k3 := 2
	r3 := false

	resu1 := containsNearbyDuplicate(nums1, k1)
	resu2 := containsNearbyDuplicate(nums2, k2)
	resu3 := containsNearbyDuplicate(nums3, k3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
