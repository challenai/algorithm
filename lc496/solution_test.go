package solution

import (
	"testing"
)

func TestNextGreaterElement(t *testing.T) {

	nums11 := []int{4, 1, 2}
	nums12 := []int{1, 3, 4, 2}
	r1 := []int{-1, 3, -1}
	nums21 := []int{2, 4}
	nums22 := []int{1, 2, 3, 4}
	r2 := []int{3, -1}

	rsu1 := NextGreaterElement(nums11, nums12)
	rsu2 := NextGreaterElement(nums21, nums22)

	for i := 0; i < len(r1); i++ {
		if rsu1[i] != r1[i] {
			t.Fail()
		}
	}

	for i := 0; i < len(r2); i++ {
		if rsu2[i] != r2[i] {
			t.Fail()
		}
	}
}
