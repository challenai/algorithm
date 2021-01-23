package solution

import (
	"testing"
)

func TestFindMaximizedCapital(t *testing.T) {

	k := 2
	W := 1
	Profits := []int{1, 2, 3}
	Captical := []int{0, 1, 1}
	r := 6
	rsu := findMaximizedCapital(k, W, Profits, Captical)
	println(rsu)

	if rsu != r {
		t.Fail()
	}
}
