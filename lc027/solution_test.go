package solution

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {

	nums1 := []int{3, 2, 2, 3}
	val1 := 3
	r1 := 2
	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := 2
	r2 := 5
	resu1 := removeElement(nums1, val1)
	resu2 := removeElement(nums2, val2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
