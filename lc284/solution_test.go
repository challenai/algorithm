package solution

import (
	"testing"
)

func TestPeekIt(t *testing.T) {
	it := Constructor(MockIterator())

	if it.Next() != 1 || it.Peek() != 2 || it.Next() != 2 || it.Next() != 3 || it.HasNext() {
		t.Fail()
	}
}
