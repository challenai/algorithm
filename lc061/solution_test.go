package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	head1 := &ListNode{
		Val: 1,
	}
	head1.Next = &ListNode{
		Val: 2,
	}
	head1.Next.Next = &ListNode{
		Val: 3,
	}
	head1.Next.Next.Next = &ListNode{
		Val: 4,
	}
	head1.Next.Next.Next.Next = &ListNode{
		Val: 5,
	}
	k1 := 2
	r1 := []int{4, 5, 1, 2, 3}

	head2 := &ListNode{
		Val: 0,
	}
	head2.Next = &ListNode{
		Val: 1,
	}
	head2.Next.Next = &ListNode{
		Val: 2,
	}
	k2 := 4
	r2 := []int{2, 0, 1}
	resu1 := rotate(head1, k1)
	resu2 := rotate(head2, k2)

	pass := true
	for i := 0; i < len(r1); i++ {
		if r1[i] != resu1.Val {
			pass = false
			break
		}
		resu1 = resu1.Next
	}
	for i := 0; i < len(r2); i++ {
		if r2[i] != resu2.Val {
			pass = false
			break
		}
		resu2 = resu2.Next
	}

	if !pass {
		t.Fail()
	}
}
