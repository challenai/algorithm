package solution

import (
	"testing"
)

func TestConstructRectangle(t *testing.T) {

	target := 195
	r := []int{13, 15}
	rsu := ConstructRectangle(target)

	if rsu[0] != r[0] {
		t.Fail()
	}
}
