package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	head.Next = &ListNode{
		Val: 2,
	}
	head.Next.Next = &ListNode{
		Val: 3,
	}
	head.Next.Next.Next = &ListNode{
		Val: 4,
	}
	head.Next.Next.Next.Next = &ListNode{
		Val: 5,
	}
	head.Next.Next.Next.Next.Next = &ListNode{
		Val: 6,
	}
	m := 2
	n := 4
	r := []int{1, 4, 3, 2, 5, 6}

	resu := reverseBetween(head, m, n)

	for i := 0; i < len(r); i++ {
		if r[i] != resu.Val {
			t.Fail()
			break
		}
		resu = resu.Next
	}
}
