package solution

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
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
	r1 := list

	list2 := &ListNode{
		Val: 5,
	}
	list2.Next = &ListNode{
		Val: 6,
	}
	list2.Next.Next = &ListNode{
		Val: 4,
	}
	var r2 *ListNode
	r2 = nil

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
	r3 := list3.Next

	rsu1 := detectCycle(list)
	rsu2 := detectCycle(list2)
	rsu3 := detectCycle(list3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
