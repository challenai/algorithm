package solution

import (
	"testing"
)

func TestCountBattleships(t *testing.T) {
	board := [][]byte{
		{'X', '.', '.', 'X'},
		{'.', 'X', '.', 'X'},
		{'.', 'X', '.', 'X'},
		{'.', 'X', '.', 'X'},
		{'.', '.', '.', 'X'},
	}
	r := 3
	rsu := countBattleships(board)

	if rsu != r {
		t.Fail()
	}
}
