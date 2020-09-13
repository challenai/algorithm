package solution

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	num1 := &ListNode{
		Val: 2,
	}
	num1.Next = &ListNode{
		Val: 4,
	}
	num1.Next.Next = &ListNode{
		Val: 3,
	}

	num2 := &ListNode{
		Val: 5,
	}
	num2.Next = &ListNode{
		Val: 6,
	}
	num2.Next.Next = &ListNode{
		Val: 4,
	}

	resu := addTwoNumbers(num1, num2)

	sum := 0
	ptr := resu
	for ptr != nil {
		sum += ptr.Val
		ptr = ptr.Next
	}

	if sum != 15 {
		t.Fail()
	}
}
