package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var ptr, temp, resu *ListNode
	resu = &ListNode{
		Next: head,
	}
	ptr = resu

	for ptr != nil && ptr.Next != nil && ptr.Next.Next != nil {
		temp = ptr.Next.Next.Next
		ptr.Next.Next.Next = ptr.Next
		ptr.Next = ptr.Next.Next
		ptr.Next.Next.Next = temp

		ptr = ptr.Next.Next
	}
	return resu.Next
}
