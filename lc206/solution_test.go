package solution

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	num1 := &ListNode{
		Val: 1,
	}
	num1.Next = &ListNode{
		Val: 2,
	}
	num1.Next.Next = &ListNode{
		Val: 3,
	}
	num1.Next.Next.Next = &ListNode{
		Val: 4,
	}
	num1.Next.Next.Next.Next = &ListNode{
		Val: 5,
	}
	num1.Next.Next.Next.Next.Next = &ListNode{
		Val: 6,
	}

	resu := reverseList(num1)
	last := 6
	for resu != nil {
		if resu.Val != last {
			t.Fail()
		}
		last--
		resu = resu.Next
	}
}
