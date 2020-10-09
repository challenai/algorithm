package solution

import (
	"testing"
)

func TestMyStack(t *testing.T) {
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	if myStack.Top() != 2 {
		t.Fail()
	} // return 2
	if myStack.Pop() != 2 {
		t.Fail()
	} // return 2
	if myStack.Empty() {
		t.Fail()
	} // return False
}
