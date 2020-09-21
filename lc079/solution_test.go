package solution

import (
	"testing"
)

func TestExist(t *testing.T) {

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	const (
		word1 = "ABCCED"
		r1    = true
		word2 = "SEE"
		r2    = true
		word3 = "ABCB"
		r3    = false
	)

	resu1 := exist(board, word1)
	resu2 := exist(board, word2)
	resu3 := exist(board, word3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
