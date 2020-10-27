package solution

import (
	"testing"
)

func TestWiggleMaxLength(t *testing.T) {
	nums1 := []int{1, 7, 4, 9, 2, 5}
	r1 := 6
	nums2 := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	r2 := 7
	nums3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r3 := 2

	rsu1 := wiggleMaxLength(nums1)
	rsu2 := wiggleMaxLength(nums2)
	rsu3 := wiggleMaxLength(nums3)

	println(rsu1)
	println(rsu2)
	println(rsu3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
