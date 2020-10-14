package solution

import (
	"testing"
)

func TestCountSmaller(t *testing.T) {
	nums := []int{5, 2, 6, 1}
	r := []int{2, 1, 1, 0}

	resu := countSmaller(nums)

	if len(resu) != len(r) {
		t.Fail()
	}
	for i := 0; i < len(resu); i++ {
		if resu[i] != r[i] {
			t.Fail()
		}
	}
}
