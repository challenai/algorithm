package solution

import (
	"testing"
)

func TestRemoveDuplicate2(t *testing.T) {
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
		Val: 3,
	}
	head1.Next.Next.Next.Next = &ListNode{
		Val: 4,
	}
	head1.Next.Next.Next.Next.Next = &ListNode{
		Val: 4,
	}
	head1.Next.Next.Next.Next.Next.Next = &ListNode{
		Val: 5,
	}
	r1 := []int{1, 2, 5}

	head2 := &ListNode{
		Val: 1,
	}
	head2.Next = &ListNode{
		Val: 1,
	}
	head2.Next.Next = &ListNode{
		Val: 1,
	}
	head2.Next.Next.Next = &ListNode{
		Val: 2,
	}
	head2.Next.Next.Next.Next = &ListNode{
		Val: 3,
	}
	r2 := []int{2, 3}
	resu1 := removeDuplicate2(head1)
	resu2 := removeDuplicate2(head2)

	for i := 0; i < len(r1); i++ {
		if r1[i] != resu1.Val {
			t.Fail()
			break
		}
		resu1 = resu1.Next
	}
	for i := 0; i < len(r2); i++ {
		if r2[i] != resu2.Val {
			t.Fail()
			break
		}
		resu2 = resu2.Next
	}
}
