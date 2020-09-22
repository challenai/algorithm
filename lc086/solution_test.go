package solution

import (
	"testing"
)

func TestPartition(t *testing.T) {
	head1 := &ListNode{
		Val: 1,
	}
	head1.Next = &ListNode{
		Val: 4,
	}
	head1.Next.Next = &ListNode{
		Val: 3,
	}
	head1.Next.Next.Next = &ListNode{
		Val: 2,
	}
	head1.Next.Next.Next.Next = &ListNode{
		Val: 5,
	}
	head1.Next.Next.Next.Next.Next = &ListNode{
		Val: 2,
	}
	x := 3
	r := []int{1, 2, 2, 4, 3, 5}

	resu1 := partition(head1, x)
	for i := 0; i < len(r); i++ {
		if r[i] != resu1.Val {
			t.Fail()
			break
		}
		resu1 = resu1.Next
	}
}
