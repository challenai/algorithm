package solution

import (
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k1 := 3
	r1 := []int{3, 3, 5, 5, 6, 7}
	nums2 := []int{9, 11}
	k2 := 2
	r2 := []int{11}
	nums3 := []int{1, -1}
	k3 := 1
	r3 := []int{1, -1}

	resu1 := maxSlidingWindow(nums1, k1)
	resu2 := maxSlidingWindow(nums2, k2)
	resu3 := maxSlidingWindow(nums3, k3)

	if len(r1) != len(resu1) {
		t.Fail()
	}
	for i := 0; i < len(r1); i++ {
		if r1[i] != resu1[i] {
			t.Fail()
		}
	}
	if len(r2) != len(resu2) {
		t.Fail()
	}
	for i := 0; i < len(r2); i++ {
		if resu2[i] != r2[i] {
			t.Fail()
		}
	}
	if len(r3) != len(resu3) {
		t.Fail()
	}
	for i := 0; i < len(r3); i++ {
		if resu3[i] != r3[i] {
			t.Fail()
		}
	}
}
