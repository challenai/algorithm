package solution

import (
	"testing"
)

func TestFindMinStep(t *testing.T) {
	board1 := "WRRBBW"
	hand1 := "RB"
	r1 := -1
	board2 := "WWRRBBWW"
	hand2 := "WRBRW"
	r2 := 2
	board3 := "G"
	hand3 := "GGGGG"
	r3 := 2
	board4 := "RBYYBBRRB"
	hand4 := "YRBGB"
	r4 := 3

	rsu1 := FindMinStep(board1, hand1)
	rsu2 := FindMinStep(board2, hand2)
	rsu3 := FindMinStep(board3, hand3)
	rsu4 := FindMinStep(board4, hand4)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 || rsu4 != r4 {
		t.Fail()
	}
}
