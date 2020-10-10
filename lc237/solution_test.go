package solution

import (
	"testing"
)

func TestX(t *testing.T) {
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
	r1 := []int{1, 3, 4, 5}

	num2 := &ListNode{
		Val: 5,
	}
	num2.Next = &ListNode{
		Val: 6,
	}
	num2.Next.Next = &ListNode{
		Val: 7,
	}
	r2 := []int{6, 7}

	deleteNode(num1.Next)
	deleteNode(num2)

	ptr := num1
	for i := 0; i < len(r1); i++ {
		if r1[i] != ptr.Val {
			t.Fail()
		}
		ptr = ptr.Next
	}
	ptr = num2
	for i := 0; i < len(r2); i++ {
		if r2[i] != ptr.Val {
			t.Fail()
		}
		ptr = ptr.Next
	}
}
