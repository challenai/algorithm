package solution

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(3)
	minStack.GetMin()
	minStack.Pop()

	if minStack.Top() != 0 {
		t.Fail()
	}

	if minStack.GetMin() != -2 {
		t.Fail()
	}
}
