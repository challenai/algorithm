package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(l1, l2 *ListNode) *ListNode {
	var temp, ptr *ListNode
	ptr = l1
	for ptr != nil && ptr.Next != nil {
		ptr = ptr.Next
	}
	temp = ptr
	ptr = l2
	for ptr != nil && ptr.Next != nil {
		ptr = ptr.Next
	}
	if ptr != temp {
		return nil
	}
	ptr.Next = l1
	ptr = l2.Next
	temp = l2.Next.Next
	for ptr != temp {
		ptr = ptr.Next
		temp = temp.Next.Next
	}
	ptr = l2
	for ptr != temp {
		ptr = ptr.Next
		temp = temp.Next
	}
	return ptr
}
