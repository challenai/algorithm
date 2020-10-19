package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	ptr := head
	head1 := &ListNode{}
	ptr1 := head1
	for ptr != nil && ptr.Next != nil {
		ptr1 = ptr.Next
		ptr1.Next = nil
		ptr.Next = ptr.Next.Next
		ptr = ptr.Next
	}
	ptr.Next = head1.Next
	ptr = head
	for ptr != nil {
		print(ptr.Val, " ")
		ptr = ptr.Next
	}
	return head
}
