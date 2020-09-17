package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	l1 := &ListNode{1, nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{5, nil}
	l2 := &ListNode{1, nil}
	l2.Next = &ListNode{3, nil}
	l2.Next.Next = &ListNode{4, nil}
	l3 := &ListNode{2, nil}
	l3.Next = &ListNode{6, nil}
	r := []int{1, 1, 2, 3, 4, 4, 5, 6}

	resu := mergeKLists([]*ListNode{l1, l2, l3})

	pass := true
	for i := 0; i < len(r); i++ {
		if resu.Val != r[i] {
			pass = false
			break
		}
		resu = resu.Next
	}
	if resu != nil {
		pass = false
	}
	if !pass {
		t.Fail()
	}
}
