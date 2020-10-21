package solution

import (
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5}
	r1 := true
	nums2 := []int{5, 4, 3, 2, 1}
	r2 := false
	nums3 := []int{3, 3, 25, 1, 34, 5, 6, 63, 6, 7}
	r3 := true

	rsu1 := increasingTriplet(nums1)
	rsu2 := increasingTriplet(nums2)
	rsu3 := increasingTriplet(nums3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
