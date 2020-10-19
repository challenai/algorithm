package solution

import (
	"testing"
)

func TestOddEvenList(t *testing.T) {
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
	list.Next.Next.Next.Next = &ListNode{
		Val: 5,
	}
	list.Next.Next.Next.Next.Next = &ListNode{
		Val: 6,
	}
	r := []int{1, 3, 5, 2, 4, 6}

	rsu := oddEvenList(list)

	ptr := rsu
	for i := 0; i < len(r); i++ {
		if r[i] != ptr.Val {
			t.Fail()
		}
		ptr = ptr.Next
	}
}
