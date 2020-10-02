package solution

import (
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	head1 := &Node{
		Val: 7,
	}
	head1.Next = &Node{
		Val: 13,
	}
	head1.Next.Next = &Node{
		Val: 11,
	}
	head1.Next.Next.Next = &Node{
		Val: 10,
	}
	head1.Next.Next.Next.Next = &Node{
		Val: 1,
	}
	head1.Next.Random = head1
	head1.Next.Next.Random = head1.Next.Next.Next.Next
	head1.Next.Next.Next.Random = head1.Next.Next
	head1.Next.Next.Next.Next.Random = head1
	resu1 := copyRandomList(head1)

	if resu1.Next.Next.Next.Next.Random.Next.Next.Next.Random.Val != 11 {
		t.Fail()
	}
}
