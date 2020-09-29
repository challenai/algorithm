package solution

import (
	"testing"
)

func TestSolve(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			print(string(board[i][j]))
		}
		println()
	}

	if false {
		t.Fail()
	}
}
