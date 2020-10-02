package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	nums1 := []int{2, 3, -2, 4}
	r1 := 6
	nums2 := []int{-2, 0, -1}
	r2 := 0

	resu1 := maxProduct(nums1)
	resu2 := maxProduct(nums2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
