package solution

import (
	"testing"
)

func TestIsSelfCrossing(t *testing.T) {
	x1 := []int{2, 1, 1, 2}
	r1 := true
	x2 := []int{1, 2, 3, 4}
	r2 := false
	x3 := []int{1, 1, 1, 1}
	r3 := true

	rsu1 := isSelfCrossing(x1)
	rsu2 := isSelfCrossing(x2)
	rsu3 := isSelfCrossing(x3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
