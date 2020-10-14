package solution

import (
	"testing"
)

func TestGameOfLife(t *testing.T) {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	r := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}
	gameOfLife(board)
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			if r[i][j] != board[i][j] {
				t.Fail()
			}
		}
	}
}
