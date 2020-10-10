package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(num *ListNode) bool {
	// the problem is we need to use only O(1) sc, so we cant save the path

	// find mid
	// reverse second half part of list
	// compare two list
	var fast, slow, par, temp *ListNode
	slow = num
	fast = num
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	par = nil
	fast = slow.Next
	for fast != nil && fast.Next != nil {
		temp = fast.Next
		fast.Next = par
		par = fast
		fast = temp
	}
	// bug...
	// fast = par
	slow = num
	for fast != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}
