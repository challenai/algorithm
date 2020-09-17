package solution

import (
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	l := &ListNode{1, nil}
	l.Next = &ListNode{2, nil}
	l.Next.Next = &ListNode{3, nil}
	l.Next.Next.Next = &ListNode{4, nil}
	l.Next.Next.Next.Next = &ListNode{5, nil}
	l.Next.Next.Next.Next.Next = &ListNode{6, nil}
	k1 := 3
	r := []int{3, 2, 1, 6, 5, 4}
	resu := reverseKGroup(l, k1)

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
