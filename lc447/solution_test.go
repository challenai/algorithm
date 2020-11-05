package solution

import (
	"testing"
)

func TestNumberOfBoomerangs(t *testing.T) {
	points1 := [][]int{{0, 0}, {1, 0}, {2, 0}}
	r1 := 2
	points2 := [][]int{{1, 1}, {2, 2}, {3, 3}}
	r2 := 2
	points3 := [][]int{{1, 1}}
	r3 := 0

	rsu1 := numberOfBoomerangs(points1)
	rsu2 := numberOfBoomerangs(points2)
	rsu3 := numberOfBoomerangs(points3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
