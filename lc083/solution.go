package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicate(head *ListNode) *ListNode {
	var ptr, before *ListNode
	ptr = head.Next
	before = head
	for ptr != nil {
		if ptr.Val == before.Val {
			ptr = ptr.Next
			before.Next = ptr
		} else {
			before = ptr
			ptr = ptr.Next
		}
	}
	return head
}
