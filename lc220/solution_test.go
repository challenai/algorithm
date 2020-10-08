package solution

import (
	"testing"
)

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	nums1 := []int{1, 2, 3, 1}
	k1 := 3
	t1 := 0
	r1 := true
	nums2 := []int{1, 0, 1, 1}
	k2 := 1
	t2 := 2
	r2 := true
	nums3 := []int{1, 5, 9, 1, 5, 9}
	k3 := 2
	t3 := 3
	r3 := true

	resu1 := containsNearbyAlmostDuplicate(nums1, k1, t1)
	resu2 := containsNearbyAlmostDuplicate(nums2, k2, t2)
	resu3 := containsNearbyAlmostDuplicate(nums3, k3, t3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
