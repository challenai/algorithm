package solution

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	n := 2

	resu := removeNthFromEnd(head, n)

	sum := 0
	var ptr *ListNode
	ptr = resu
	for ptr != nil {
		sum += ptr.Val
		ptr = ptr.Next
	}

	if sum != 11 {
		t.Fail()
	}
}
