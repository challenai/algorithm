package solution

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
	x1 := 1
	y1 := 4
	r1 := 2

	rsu1 := hammingDistance(x1, y1)

	if rsu1 != r1 {
		t.Fail()
	}
}
