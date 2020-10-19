package solution

import (
	"testing"
)

func TestIntersect(t *testing.T) {
	nums11 := []int{1, 2, 2, 1}
	nums12 := []int{2, 2}
	r1 := []int{2, 2}
	nums21 := []int{4, 9, 5}
	nums22 := []int{9, 4, 9, 8, 4}
	r2 := []int{4, 9}
	rsu1 := intersect(nums11, nums12)
	rsu2 := intersect(nums21, nums22)

	if len(rsu1) != len(r1) || len(rsu2) != len(r2) {
		t.Fail()
	}
}
