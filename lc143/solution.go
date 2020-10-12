package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	dq := []*ListNode{}
	ptr := head
	for ptr != nil {
		dq = append(dq, ptr)
		ptr = ptr.Next
	}
	head.Next = dq[len(dq)-1]
	dq = dq[1 : len(dq)-1]
	ptr = head.Next
	for len(dq) > 0 {
		ptr.Next = dq[0]
		dq = dq[1:]
		ptr = ptr.Next
		if len(dq) > 0 {
			ptr.Next = dq[len(dq)-1]
			dq = dq[:len(dq)-1]
			ptr = ptr.Next
		}
	}
	ptr.Next = nil
}
