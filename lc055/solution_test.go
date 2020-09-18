package solution

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	nums1 := []int{2, 3, 1, 1, 4}
	r1 := true
	nums2 := []int{3, 2, 1, 0, 4}
	r2 := false

	resu1 := canJump(nums1)
	resu2 := canJump(nums2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
