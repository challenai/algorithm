package solution

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	s := "abbaca"
	r := "ca"

	rsu := removeDuplicates(s)

	if rsu != r {
		t.Fail()
	}
}
