package solution

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	list := &ListNode{
		Val: 7,
	}
	list.Next = &ListNode{
		Val: 2,
	}
	list.Next.Next = &ListNode{
		Val: 4,
	}
	list.Next.Next.Next = &ListNode{
		Val: 3,
	}

	list2 := &ListNode{
		Val: 5,
	}
	list2.Next = &ListNode{
		Val: 6,
	}
	list2.Next.Next = &ListNode{
		Val: 4,
	}

	r := []int{7, 8, 0, 7}

	rsu := addTwoNumbers(list, list2)

	for i := 0; i < len(r); i++ {
		if rsu.Val != r[i] {
			t.Fail()
		}
		rsu = rsu.Next
	}
}
