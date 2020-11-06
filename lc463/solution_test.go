package solution

import (
	"testing"
)

func TestIslandPerimeter(t *testing.T) {
	grid1 := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	r1 := 16
	grid2 := [][]int{{1, 0}}
	r2 := 4

	rsu1 := islandPerimeter(grid1)
	rsu2 := islandPerimeter(grid2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
