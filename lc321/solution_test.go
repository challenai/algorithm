package solution

import (
	"testing"
)

func TestMaxNumber(t *testing.T) {
	nums11 := []int{3, 4, 6, 5}
	nums12 := []int{9, 1, 2, 5, 8, 3}
	k1 := 5
	r1 := []int{9, 8, 6, 5, 3}
	nums21 := []int{6, 7}
	nums22 := []int{6, 0, 4}
	k2 := 5
	r2 := []int{6, 7, 6, 0, 4}
	nums31 := []int{3, 9}
	nums32 := []int{8, 9}
	k3 := 3
	r3 := []int{9, 8, 9}

	rsu1 := maxNumber(nums11, nums12, k1)
	rsu2 := maxNumber(nums21, nums22, k2)
	rsu3 := maxNumber(nums31, nums32, k3)

	if len(rsu1) != len(r1) || len(rsu2) != len(r2) || len(rsu3) != len(r3) {
		t.Fail()
	}
}
