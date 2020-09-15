package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var ptr *ListNode
	ptr = head
	st := []*ListNode{}
	for ptr != nil {
		st = append(st, ptr)
		ptr = ptr.Next
	}
	if n > len(st) {
		return head
	}
	if n == len(st) {
		return head.Next
	}
	st[len(st)-n-1].Next = st[len(st)-n+1]
	st[len(st)-n] = nil
	return head
}
