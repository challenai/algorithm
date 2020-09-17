package solution

import (
	"testing"
)

func TestSearch(t *testing.T) {
	nums1 := []int{4, 5, 6, 7, 0, 1, 2, 3}
	target1 := 0
	r1 := 4
	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	target2 := 3
	r2 := -1

	resu1 := search(nums1, target1)
	resu2 := search(nums2, target2)
    println(resu1)
    println(resu2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
