package solution

import (
	"testing"
)

func TestFindMinArrowShots(t *testing.T) {
	points1 := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	r1 := 2
	points2 := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	r2 := 4
	points3 := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	r3 := 2
	points4 := [][]int{{1, 2}}
	r4 := 1
	points5 := [][]int{{2, 3}, {2, 3}}
	r5 := 1

	rsu1 := findMinArrowShots(points1)
	rsu2 := findMinArrowShots(points2)
	rsu3 := findMinArrowShots(points3)
	rsu4 := findMinArrowShots(points4)
	rsu5 := findMinArrowShots(points5)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 || rsu4 != r4 || rsu5 != r5 {
		t.Fail()
	}
}
