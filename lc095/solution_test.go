package solution

import (
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	n := 3
	r := 5

	resu := generateTrees(n)
	if len(resu) != r {
		t.Fail()
	}
}
