package solution

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{1, nil}
	l1.Next = &ListNode{2, nil}
	l1.Next.Next = &ListNode{5, nil}
	l1.Next.Next.Next = &ListNode{7, nil}
	l2 := &ListNode{1, nil}
	l2.Next = &ListNode{3, nil}
	l2.Next.Next = &ListNode{4, nil}
	resu := mergeTwoLists(l1, l2)

	var temp int
	var ptr *ListNode
	var pass bool
	ptr = resu
	pass = true
	temp = ptr.Val
	for ptr != nil {
		// println(ptr.Val)
		if ptr.Val < temp {
			pass = false
			break
		}
		temp = ptr.Val
		ptr = ptr.Next
	}

	if !pass {
		t.Fail()
	}
}
