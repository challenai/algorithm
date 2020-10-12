package solution

import (
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	list := &ListNode{
		Val: 3,
	}
	list.Next = &ListNode{
		Val: 1,
	}
	list.Next.Next = &ListNode{
		Val: 6,
	}
	list.Next.Next.Next = &ListNode{
		Val: 5,
	}
	list.Next.Next.Next.Next = &ListNode{
		Val: 2,
	}
	list.Next.Next.Next.Next.Next = &ListNode{
		Val: 4,
	}

	resu := insertionSortList(list)
	par := resu
	ptr := resu.Next
	for ptr != nil {
		if par.Val > ptr.Val {
			t.Fail()
		}
		par = ptr
		ptr = ptr.Next
	}
}
