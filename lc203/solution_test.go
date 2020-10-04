package solution

import (
	"testing"
)

func TestRemoveElements(t *testing.T) {
	num1 := &ListNode{
		Val: 1,
	}
	num1.Next = &ListNode{
		Val: 2,
	}
	num1.Next.Next = &ListNode{
		Val: 2,
	}
	num1.Next.Next.Next = &ListNode{
		Val: 3,
	}
	num1.Next.Next.Next.Next = &ListNode{
		Val: 2,
	}
	num1.Next.Next.Next.Next.Next = &ListNode{
		Val: 3,
	}
	val := 2
	r := []int{1, 3, 3}
	rsu := removeElements(num1, val)
	for i := 0; i < len(r); i++ {
		if r[i] != rsu.Val {
			t.Fail()
		}
		rsu = rsu.Next
	}
}
