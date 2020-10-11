package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	nums1 := []int{3, 0, 1}
	r1 := 2
	nums2 := []int{0, 1}
	r2 := 2
	nums3 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	r3 := 8

	rsu1 := missingNumber(nums1)
	rsu2 := missingNumber(nums2)
	rsu3 := missingNumber(nums3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
