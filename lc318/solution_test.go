package solution

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	words1 := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	r1 := 16
	words2 := []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	r2 := 4
	words3 := []string{"a", "aa", "aaa", "aaaa"}
	r3 := 0

	rsu1 := maxProduct(words1)
	rsu2 := maxProduct(words2)
	rsu3 := maxProduct(words3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
