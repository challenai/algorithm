package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, target int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := &ListNode{
		Val: 0,
	}
	var ptr, ptrTemp *ListNode
	ptrTemp = temp
	ptr = head
	for ptr != nil && ptr.Next != nil {
		if ptr.Next.Val >= target {
			ptrTemp.Next = ptr.Next
			ptrTemp = ptrTemp.Next
			ptr.Next = ptr.Next.Next
		} else {
			ptr = ptr.Next
		}
	}

	ptrTemp.Next = nil
	ptr.Next = temp.Next
	ptr = head
	// for ptr != nil {
	// 	print(ptr.Val, " ")
	// 	ptr = ptr.Next
	// }
	// println()
	return head
}
