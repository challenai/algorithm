package solution

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	r1 := 1
	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	r2 := 3

	rsu1 := numIslands(grid1)
	rsu2 := numIslands(grid2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
