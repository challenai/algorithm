package solution

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	list := &ListNode{
		Val: 1,
	}
	list.Next = &ListNode{
		Val: 2,
	}
	list.Next.Next = &ListNode{
		Val: 3,
	}
	list.Next.Next.Next = &ListNode{
		Val: 4,
	}
	list.Next.Next.Next = &ListNode{
		Val: 5,
	}
	list.Next.Next.Next.Next = &ListNode{
		Val: 6,
	}
	list.Next.Next.Next.Next.Next = list
	r1 := true

	list2 := &ListNode{
		Val: 5,
	}
	list2.Next = &ListNode{
		Val: 6,
	}
	list2.Next.Next = &ListNode{
		Val: 4,
	}
	r2 := false

	list3 := &ListNode{
		Val: 5,
	}
	list3.Next = &ListNode{
		Val: 6,
	}
	list3.Next.Next = &ListNode{
		Val: 4,
	}
	list3.Next.Next.Next = list3.Next
	r3 := true

	rsu1 := hasCycle(list)
	rsu2 := hasCycle(list2)
	rsu3 := hasCycle(list3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
