package solution

import (
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	l1 := &ListNode{1, nil}
	l1.Next = &ListNode{2, nil}

	l2 := &ListNode{1, nil}
	l2.Next = &ListNode{3, nil}
	l2.Next.Next = &ListNode{4, nil}
	l2.Next.Next.Next = &ListNode{6, nil}
	l2.Next.Next.Next.Next = &ListNode{8, nil}

	l1.Next.Next = l2.Next.Next.Next

	resu := getIntersectionNode(l1, l2)

	if resu == nil || resu.Val != 6 {
		t.Fail()
	}
}
