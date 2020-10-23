package solution

import (
	"testing"
)

func TestCanMeasureWater(t *testing.T) {
	x := 3
	y := 5
	z := 4
	r := true
	x1 := 2
	y1 := 6
	z1 := 5
	r1 := false

	rsu := canMeasureWater(x, y, z)
	rsu1 := canMeasureWater(x1, y1, z1)

	if rsu1 != r1 || rsu != r {
		t.Fail()
	}
}
