package solution

import (
	"testing"
)

func TestIsBoomerang(t *testing.T) {
	points1 := [][]int{{1, 1}, {2, 3}, {3, 2}}
	r1 := true
	points2 := [][]int{{1, 1}, {2, 2}, {3, 3}}
	r2 := false
	points3 := [][]int{{0, 1}, {0, 2}, {0, 3}}
	r3 := false
	points4 := [][]int{{1, 0}, {2, 0}, {3, 0}}
	r4 := false

	rsu1 := isBoomerang(points1)
	rsu2 := isBoomerang(points2)
	rsu3 := isBoomerang(points3)
	rsu4 := isBoomerang(points4)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 || rsu4 != r4 {
		t.Fail()
	}
}
