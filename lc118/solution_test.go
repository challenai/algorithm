package solution

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	num := 5
	r := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}

	resu := generate(num)

	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			if r[i][j] != resu[i][j] {
				t.Fail()
			}
		}
	}
}
