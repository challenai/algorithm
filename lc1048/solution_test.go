package solution

import (
	"testing"
)

func TestLongestStrChain(t *testing.T) {
	words := []string{"a", "b", "ba", "bca", "bda", "bdca"}
	r := 4

	rsu := longestStrChain(words)

	if rsu != r {
		t.Fail()
	}
}
