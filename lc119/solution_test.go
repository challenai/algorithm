package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	rowIndex1 := 5
	r1 := []int{1, 5, 10, 10, 5, 1}
	rowIndex2 := 1
	r2 := []int{1, 1}

	resu1 := getRow(rowIndex1)
	resu2 := getRow(rowIndex2)

	if len(resu1) != len(r1) || len(resu2) != len(r2) {
		t.Fail()
	}
}
