package solution

import (
	"testing"
)

func TestSearch(t *testing.T) {

	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target1 := 0
	r1 := true
	target2 := 3
	r2 := false
	resu1 := search(nums, target1)
	resu2 := search(nums, target2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
