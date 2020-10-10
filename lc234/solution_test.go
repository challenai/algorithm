package solution

import (
	"testing"
)

func TestIsParalindrome(t *testing.T) {
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
		Val: 3,
	}
	num1.Next.Next.Next.Next.Next = &ListNode{
		Val: 2,
	}
	num1.Next.Next.Next.Next.Next.Next = &ListNode{
		Val: 1,
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
	resu := isPalindrome(num1)
	resu2 := isPalindrome(num2)

	if !resu || resu2 {
		t.Fail()
	}
}
