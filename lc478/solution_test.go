package solution

import (
	"testing"
)

func TestRandPoint(t *testing.T) {
	radius := 10.0
	x_center := 5.0
	y_center := -7.5
	s := Constructor(radius, x_center, y_center)
	p := s.RandPoint()
	println(p[0], p[1])
	p = s.RandPoint()
	println(p[0], p[1])

	t.Fail()
}
