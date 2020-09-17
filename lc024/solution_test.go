package solution

import (
	"testing"
)

func TestSwapPairs(t *testing.T) {
	l := &ListNode{1, nil}
	l.Next = &ListNode{2, nil}
	l.Next.Next = &ListNode{3, nil}
	l.Next.Next.Next = &ListNode{4, nil}
	r := []int{2, 1, 4, 3}
	resu := swapPairs(l)

	pass := true
	for i := 0; i < len(r); i++ {
		if r[i] != resu.Val {
			pass = false
			break
		}
		resu = resu.Next
	}

	if !pass {
		t.Fail()
	}
}
