package solution

import (
	"testing"
)

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	r := 6

	resu := trap(height)

	if resu != r {
		t.Fail()
	}
}
