package solution

import (
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	r := []int{5, 6}

	rsu := findDisappearedNumbers(nums)

	if len(rsu) != len(r) {
		t.Fail()
	}
	for i := 0; i < len(r); i++ {
		if r[i] != rsu[i] {
			t.Fail()
		}
	}
}
