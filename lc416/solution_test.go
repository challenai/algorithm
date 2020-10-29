package solution

import (
	"testing"
)

func TestCanPartition(t *testing.T) {
	nums1 := []int{1, 2, 11, 5, 12, 9}
	r1 := true
	nums2 := []int{1, 2, 3, 5}
	r2 := false

	rsu1 := canPartition(nums1)
	rsu2 := canPartition(nums2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
