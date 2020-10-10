package solution

import (
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	nums1 := []int{1, 2, 3, 4}
	r1 := []int{24, 12, 8, 6}
	nums2 := []int{2, -4, 5, 6, -7, 0, 8, -3}
	r2 := []int{0, 0, 0, 0, 0, -40320, 0, 0}

	resu1 := productExceptSelf(nums1)
	resu2 := productExceptSelf(nums2)

	if len(resu1) != len(r1) {
		t.Fail()
	}
	if len(resu2) != len(r2) {
		t.Fail()
	}
	for i := 0; i < len(r1); i++ {
		if r1[i] != resu1[i] {
			t.Fail()
		}
	}
	for i := 0; i < len(r2); i++ {
		if r2[i] != resu2[i] {
			t.Fail()
		}
	}

}
