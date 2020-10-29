package solution

import (
	"testing"
)

func TestThirdMax(t *testing.T) {
	nums1 := []int{3, 2, 1}
	r1 := 1
	nums2 := []int{1, 2}
	r2 := 2
	nums3 := []int{2, 2, 3, 1}
	r3 := 1

	rsu1 := thirdMax(nums1)
	rsu2 := thirdMax(nums2)
	rsu3 := thirdMax(nums3)
	println(rsu1)
	println(rsu2)
	println(rsu3)

	if r1 != rsu1 || r2 != rsu2 || r3 != rsu3 {
		t.Fail()
	}
}
