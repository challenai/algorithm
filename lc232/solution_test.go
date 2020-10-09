package solution

import (
	"testing"
)

func TestMyQueue(t *testing.T) {
	myQueue := Constructor()
	myQueue.Push(1) // queue is: [1]
	myQueue.Push(2) // queue is: [1, 2] (leftmost is front of the queue)
	if myQueue.Peek() != 1 {
		t.Fail()
	} // return 1
	if myQueue.Pop() != 1 {
		t.Fail()
	} // return 1, queue is [2]
	if myQueue.Empty() {
		t.Fail()
	} // return false
	myQueue.Pop()
	if !myQueue.Empty() {
		t.Fail()
	}
}
